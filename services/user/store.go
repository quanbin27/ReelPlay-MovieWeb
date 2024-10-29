package user

import (
	"errors"
	"fmt"
	"github.com/quanbin27/ReelPlay/services/auth"
	"github.com/quanbin27/ReelPlay/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db}
}

func (s *Store) UpdateUserPassword(userID int, newPassword string) error {
	// Mã hóa mật khẩu mới
	hashedPassword, err := auth.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Cập nhật mật khẩu đã mã hóa
	result := s.db.Model(&types.User{}).Where("id = ?", userID).Update("password", hashedPassword)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
func (s *Store) UpdateUserInfo(userID int, updatedData map[string]interface{}) error {
	// Chỉ cho phép cập nhật firstName, lastName và email
	allowedFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"role_id":    true,
	}

	for key := range updatedData {
		if !allowedFields[key] {
			delete(updatedData, key) // Xóa các trường không hợp lệ
		}
	}

	// Cập nhật thông tin người dùng
	result := s.db.Model(&types.User{}).Where("id = ?", userID).Updates(updatedData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *Store) UpdateInfo(userID int, updatedData map[string]interface{}) error {
	// Chỉ cho phép cập nhật firstName, lastName và email
	allowedFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"email":      true,
	}

	for key := range updatedData {
		if !allowedFields[key] {
			delete(updatedData, key) // Xóa các trường không hợp lệ
		}
	}

	// Cập nhật thông tin người dùng
	result := s.db.Model(&types.User{}).Where("id = ?", userID).Updates(updatedData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *Store) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	result := store.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (s *Store) UpdatePassword(userID int, newPassword string) error {
	return s.db.Model(&types.User{}).Where("id = ?", userID).Update("password", newPassword).Error
}
func (s *Store) CountUsers() (int, error) {
	var count int64
	if err := s.db.Model(&types.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

// Trong file store.go hoặc tương tự
func (s *Store) UnlockUser(userID int) error {
	var user types.User
	if err := s.db.Unscoped().First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Đánh dấu là chưa xóa
	if err := s.db.Unscoped().Model(&user).Update("deleted_at", nil).Error; err != nil {
		return err
	}

	return nil
}

func (store *Store) GetUserByID(id int) (*types.User, error) {
	var user types.User
	result := store.db.Unscoped().Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (store *Store) CreateUser(user *types.User) error {
	result := store.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *Store) DeleteUserSoft(userID int) error {
	// Xóa mềm người dùng với GORM (sử dụng gorm.DeletedAt)
	result := s.db.Where("id = ?", userID).Delete(&types.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (s *Store) SearchUsers(keyword string, page, limit int) ([]types.User, int64, error) {
	var users []types.User
	var totalRecords int64

	// Tạo truy vấn tìm kiếm người dùng
	query := s.db.Unscoped().Model(&types.User{}).Preload("Role") // Preload Role để lấy thông tin role của người dùng

	// Nếu có keyword, thêm điều kiện tìm kiếm theo FirstName, LastName hoặc Email
	if keyword != "" {
		keyword = "%" + keyword + "%" // Thêm ký tự '%' cho LIKE query
		query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ?", keyword, keyword, keyword)
	}

	// Tính tổng số lượng bản ghi
	err := query.Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	// Áp dụng pagination
	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, totalRecords, nil
}

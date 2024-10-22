package director

import (
	"github.com/quanbin27/ReelPlay/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db}
}
func (s *Store) CreateDirector(director *types.Director) error {
	// Tạo mới Director
	if err := s.db.Create(director).Error; err != nil {
		return err // Trả về lỗi nếu không tạo được
	}
	return nil // Trả về nil nếu thành công
}
func (s *Store) UpdateDirector(id int, director *types.Director) error {
	var existingDirector types.Director
	if err := s.db.First(&existingDirector, id).Error; err != nil {
		return err // Trả về lỗi nếu không tìm thấy
	}

	director.ID = existingDirector.ID

	if err := s.db.Model(&existingDirector).Updates(director).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteDirector(id int) error {
	// Tìm Director theo ID
	var director types.Director
	if err := s.db.First(&director, id).Error; err != nil {
		return err // Trả về lỗi nếu không tìm thấy
	}

	// Xóa Director
	if err := s.db.Delete(&director).Error; err != nil {
		return err // Trả về lỗi nếu có vấn đề khi xóa
	}
	return nil // Trả về nil nếu thành công
}
func (s *Store) GetAllDirectors() ([]types.Director, error) {
	var directors []types.Director
	// Lấy tất cả Directors
	if err := s.db.Find(&directors).Error; err != nil {
		return nil, err // Trả về lỗi nếu có vấn đề khi truy vấn
	}
	return directors, nil // Trả về danh sách Directors nếu thành công
}
func (s *Store) GetDirectorByID(id int) (*types.Director, error) {
	var director types.Director
	// Tìm Director theo ID
	if err := s.db.First(&director, id).Error; err != nil {
		return nil, err // Trả về lỗi nếu không tìm thấy
	}
	return &director, nil // Trả về director nếu thành công
}

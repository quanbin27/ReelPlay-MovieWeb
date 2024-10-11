package user

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

func (store *Store) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	result := store.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (s *Store) UpdatePassword(userID int, newPassword string) error {
	return s.db.Where("id = ?", userID).Update("password", newPassword).Error
}

func (store *Store) GetUserByID(id int) (*types.User, error) {
	var user types.User
	result := store.db.Where("id = ?", id).First(&user)
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

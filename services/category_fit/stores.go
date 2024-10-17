package category_fit

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
func (s *Store) CreateCategoryFit(fit *types.CategoryFit) error {
	if err := s.db.Create(fit).Error; err != nil {
		return err
	}
	return nil
}
func (s *Store) UpdateCategoryFit(fit *types.CategoryFit) error {
	if err := s.db.Save(fit).Error; err != nil {
		return err
	}
	return nil
}
func (s *Store) GetUserCategoryFit(userID int, categoryID int) (*types.CategoryFit, error) {
	var fit types.CategoryFit
	if err := s.db.Where("user_id = ? AND category_id = ?", userID, categoryID).First(&fit).Error; err != nil {
		return nil, err
	}
	return &fit, nil
}

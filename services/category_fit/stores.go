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
func (s *Store) GetCategoryFit(userId int) (cate1, cate2, cate3 *types.CategoryFit, err error) {
	var categories []types.CategoryFit

	result := s.db.Where("user_id = ? AND fit_rate > ?", userId, 5).Order("fit_rate DESC").Limit(3).Find(&categories)
	if result.Error != nil {
		return nil, nil, nil, result.Error
	}

	switch len(categories) {
	case 0:
		return nil, nil, nil, nil
	case 1:
		cate1 = &categories[0]
		return cate1, nil, nil, nil
	case 2:
		cate1 = &categories[0]
		cate2 = &categories[1]
		return cate1, cate2, nil, nil
	default:
		cate1 = &categories[0]
		cate2 = &categories[1]
		cate3 = &categories[2]
		return cate1, cate2, cate3, nil
	}
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

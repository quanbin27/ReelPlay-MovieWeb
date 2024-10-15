package rate

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
func (s *Store) CreateMovieRating(rating *types.Rate) error {
	if err := s.db.Create(rating).Error; err != nil {
		return err
	}
	return nil
}
func (s *Store) UpdateMovieRating(rating *types.Rate) error {
	if err := s.db.Save(rating).Error; err != nil {
		return err
	}
	return nil
}
func (s *Store) GetUserMovieRating(userID int, movieID int) (*types.Rate, error) {
	var rating types.Rate
	if err := s.db.Where("user_id = ? AND movie_id = ?", userID, movieID).First(&rating).Error; err != nil {
		return nil, err
	}
	return &rating, nil
}
func (s *Store) UpdateMovieAverageRating(movieID int) error {
	var totalRating int64
	var count int64

	if err := s.db.Model(&types.Rate{}).
		Where("movie_id = ?", movieID).
		Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return nil // Không cần cập nhật nếu không có đánh giá nào
	}

	if err := s.db.Model(&types.Rate{}).
		Where("movie_id = ?", movieID).
		Select("SUM(rate)").Scan(&totalRating).Error; err != nil {
		return err
	}

	// Tính rating trung bình
	averageRating := float64(totalRating) / float64(count)

	// Cập nhật lại rating trung bình cho movie
	if err := s.db.Model(&types.Movie{}).Where("id = ?", movieID).
		Update("rate", averageRating).Error; err != nil {
		return err
	}
	return nil
}

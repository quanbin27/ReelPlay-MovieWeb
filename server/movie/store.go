package movie

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
func (s *Store) GetMoviesWithPagination(offset, limit int) ([]types.Movie, error) {
	var movies []types.Movie
	result := s.db.Limit(limit).Offset(offset).Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}
func (s *Store) GetMovieById(id string) (types.Movie, error) {
	var movie types.Movie
	result := s.db.Where("id = ?", id).First(&movie)
	if result.Error != nil {
		return types.Movie{}, result.Error
	}
	return movie, nil
}

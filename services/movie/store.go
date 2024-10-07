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
func (s *Store) GetMoviesWithPagination(offset, limit int) ([]types.MovieResponse, error) {
	var movies []types.Movie
	result := s.db.Preload("Actor").Preload("Category").Preload("Director").Limit(limit).Offset(offset).Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	var movieResponses []types.MovieResponse
	for _, movie := range movies {
		var movieResponse types.MovieResponse
		movieResponse.ID = movie.ID
		movieResponse.Name = movie.Name
		movieResponse.Year = movie.Year
		movieResponse.NumEpisodes = movie.NumEpisodes
		movieResponse.Description = movie.Description
		movieResponse.Language = movie.Language
		movieResponse.Country = movie.Country
		movieResponse.Thumbnail = movie.Thumbnail
		movieResponse.Trailer = movie.Trailer
		movieResponse.Rate = movie.Rate
		movieResponse.IsFree = movie.IsFree

		for _, actor := range movie.Actor {
			movieResponse.Actor = append(movieResponse.Actor, actor.Name)
		}

		for _, director := range movie.Director {
			movieResponse.Director = append(movieResponse.Director, director.Name)
		}

		for _, category := range movie.Category {
			movieResponse.Category = append(movieResponse.Category, category.Name)
		}

		// Thêm movieResponse vào slice kết quả
		movieResponses = append(movieResponses, movieResponse)
	}

	return movieResponses, nil
}
func (s *Store) GetMovieById(id string) (types.Movie, error) {
	var movie types.Movie
	result := s.db.Where("id = ?", id).First(&movie)
	if result.Error != nil {
		return types.Movie{}, result.Error
	}
	return movie, nil
}

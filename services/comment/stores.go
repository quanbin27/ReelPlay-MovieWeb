package comment

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
func (s *Store) CreateComment(content string, movieID int, userID int) (*types.Comment, error) {
	comment := &types.Comment{
		Content: content,
		MovieID: movieID,
		UserID:  userID,
	}
	if err := s.db.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}
func (s *Store) GetCommentsByMovieID(movieID int) ([]types.Comment, error) {
	var comments []types.Comment
	err := s.db.Where("movie_id = ?", movieID).Order("created_at desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

package bookmark

import (
	"errors"
	"fmt"
	"github.com/quanbin27/ReelPlay/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db}
}
func (s *Store) IsBookmark(movieId int, userId int) bool {
	var bookmark types.Bookmark
	err := s.db.Where("user_id = ? AND movie_id = ?", userId, movieId).First(&bookmark).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}
func (s *Store) GetBookmarksByUser(userId int) ([]types.Bookmark, error) {
	var bookmarks []types.Bookmark
	if err := s.db.Where("user_id = ?", userId).Find(&bookmarks).Error; err != nil {
		return nil, err
	}
	return bookmarks, nil
}

func (s *Store) CreateBookmark(movieId int, userId int) error {
	var bookmark types.Bookmark
	err := s.db.Where("user_id = ? AND movie_id = ?", userId, movieId).First(&bookmark).Error
	if err == nil {
		return fmt.Errorf("bookmark already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Nếu chưa tồn tại, tạo bookmark mới
	bookmark = types.Bookmark{
		UserID:  userId,
		MovieID: movieId,
	}
	if err := s.db.Create(&bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) CancelBookmark(movieId int, userId int) error {
	if err := s.db.Where("user_id = ? AND movie_id = ?", userId, movieId).Delete(&types.Bookmark{}).Error; err != nil {
		return err
	}
	return nil
}

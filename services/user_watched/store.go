package user_watched

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
func (s *Store) UpdateWatchPosition(userId, episodeId, position int) error {
	return s.db.Model(&types.UserWatched{}).
		Where("user_id = ? AND episode_id = ?", userId, episodeId).
		Update("last_position", position).Error
}
func (s *Store) GetWatchPosition(userId, episodeId int) (int, error) {
	var userWatched types.UserWatched
	err := s.db.Where("user_id = ? AND episode_id = ?", userId, episodeId).
		First(&userWatched).Error
	if err != nil {
		return 0, err
	}
	return userWatched.LastPosition, nil
}

func (s *Store) CreateUserWatched(userId int, episodeId int) error {
	// Kiểm tra xem bản ghi đã tồn tại chưa
	var userWatched types.UserWatched
	result := s.db.Where("user_id = ? AND episode_id = ?", userId, episodeId).First(&userWatched)
	if result.Error == nil {
		return nil
	}
	if result.Error == gorm.ErrRecordNotFound {
		newUserWatched := types.UserWatched{
			UserID:       userId,
			EpisodeID:    episodeId,
			LastPosition: 0, // Vị trí mặc định
		}

		err := s.db.Create(&newUserWatched).Error
		return err
	}

	return result.Error
}

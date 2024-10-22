package episode

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
func (s *Store) GetEpisodeById(id int) (*types.Episode, error) {
	var episode types.Episode
	err := s.db.First(&episode, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &episode, err
}
func (s *Store) GetEpisodeByMovieAndEpisodeId(movieId, episodeNumber int) (*types.Episode, error) {
	var episode types.Episode
	err := s.db.Where("movie_id = ? AND episode_number = ?", movieId, episodeNumber).First(&episode).Error
	if err != nil {
		return nil, err
	}
	return &episode, nil
}
func (s *Store) CreateEpisode(episode *types.Episode) error {
	return s.db.Create(episode).Error
}
func (s *Store) UpdateEpisode(id int, episode *types.Episode) error {
	var existingEpisode types.Episode
	if err := s.db.First(&existingEpisode, id).Error; err != nil {
		return err // Nếu không tìm thấy episode, trả về lỗi
	}

	// Cập nhật thông tin cho episode
	episode.ID = existingEpisode.ID // Giữ nguyên ID
	return s.db.Model(&existingEpisode).Updates(episode).Error
}
func (s *Store) DeleteEpisode(id int) error {
	var episode types.Episode
	if err := s.db.First(&episode, id).Error; err != nil {
		return err // Nếu không tìm thấy episode, trả về lỗi
	}

	// Xóa episode
	return s.db.Delete(&episode).Error
}

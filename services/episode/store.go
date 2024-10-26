package episode

import (
	"github.com/quanbin27/ReelPlay/types"
	"gorm.io/gorm"
	"time"
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
func (s *Store) SearchEpisodes(keyword string, page int, limit int) ([]types.SearchEpisodesResponse, int64, error) {
	var episodes []types.SearchEpisodesResponse
	var total int64

	// Offset for pagination
	offset := (page - 1) * limit

	// Building the query with join to get the Movie title
	query := s.db.Table("episodes").
		Select("episodes.id, episodes.movie_id, episodes.episode_number, movies.name as movie_title, episodes.duration, episodes.updated_at").
		Joins("left join movies on movies.id = episodes.movie_id").
		Where("episodes.deleted_at IS NULL")

	// Add search condition
	if keyword != "" {
		searchKeyword := "%" + keyword + "%"
		query = query.Where("movies.name LIKE ? OR episodes.episode_number LIKE ? OR episodes.movie_id LIKE ?", searchKeyword, searchKeyword, searchKeyword)
	}

	// Get total count for pagination
	query.Count(&total)

	// Apply pagination and execute the query
	err := query.Limit(limit).Offset(offset).Scan(&episodes).Error
	if err != nil {
		return nil, 0, err
	}

	return episodes, total, nil
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
func (s *Store) UpdateEpisode(id int, episode *types.UpdateEpisodeRequest) error {

	return s.db.Model(&types.Episode{}).Where("id = ?", id).Update("source", episode.Source).Update("duration", episode.Duration).Update("updated_at", time.Now()).Error
}
func (s *Store) CountEpisodes() (int, error) {
	var count int64                                                    // Thử dùng int64 thay cho int
	if err := s.db.Table("episodes").Count(&count).Error; err != nil { // Chỉ định tên bảng trực tiếp
		return 0, err
	}
	return int(count), nil
}

func (s *Store) DeleteEpisode(id int) error {
	var episode types.Episode
	if err := s.db.First(&episode, id).Error; err != nil {
		return err // Nếu không tìm thấy episode, trả về lỗi
	}
	// Xóa episode
	return s.db.Delete(&episode).Error
}

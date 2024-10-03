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

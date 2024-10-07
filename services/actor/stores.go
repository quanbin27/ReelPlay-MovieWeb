package actor

import (
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db}
}

//func (s *Store) GetActorByMovieID(mvid int) ([]types.Actor, error) {
//	var actorids []int
//	err := s.db.
//	return actors, nil
//}

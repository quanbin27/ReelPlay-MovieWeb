package actor

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
func (s *Store) CreateActor(Actor *types.Actor) error {
	// Tạo mới Actor
	if err := s.db.Create(Actor).Error; err != nil {
		return err // Trả về lỗi nếu không tạo được
	}
	return nil // Trả về nil nếu thành công
}
func (s *Store) UpdateActor(id int, Actor *types.Actor) error {
	var existingActor types.Actor
	if err := s.db.First(&existingActor, id).Error; err != nil {
		return err // Trả về lỗi nếu không tìm thấy
	}

	Actor.ID = existingActor.ID

	if err := s.db.Model(&existingActor).Updates(Actor).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteActor(id int) error {
	// Tìm Actor theo ID
	var Actor types.Actor
	if err := s.db.First(&Actor, id).Error; err != nil {
		return err // Trả về lỗi nếu không tìm thấy
	}

	// Xóa Actor
	if err := s.db.Delete(&Actor).Error; err != nil {
		return err // Trả về lỗi nếu có vấn đề khi xóa
	}
	return nil // Trả về nil nếu thành công
}
func (s *Store) GetAllActors() ([]types.Actor, error) {
	var Actors []types.Actor
	// Lấy tất cả Actors
	if err := s.db.Find(&Actors).Error; err != nil {
		return nil, err // Trả về lỗi nếu có vấn đề khi truy vấn
	}
	return Actors, nil // Trả về danh sách Actors nếu thành công
}
func (s *Store) GetActorByID(id int) (*types.Actor, error) {
	var Actor types.Actor
	// Tìm Actor theo ID
	if err := s.db.First(&Actor, id).Error; err != nil {
		return nil, err // Trả về lỗi nếu không tìm thấy
	}
	return &Actor, nil // Trả về Actor nếu thành công
}

//func (s *Store) GetActorByMovieID(mvid int) ([]types.Actor, error) {
//	var actorids []int
//	err := s.db.
//	return actors, nil
//}

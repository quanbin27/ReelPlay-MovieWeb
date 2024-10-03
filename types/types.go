package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int64) (*User, error)
	CreateUser(user *User) error
}
type MovieStore interface {
	GetMoviesWithPagination(offset, limit int) ([]Movie, error)
}
type EpisodeStore interface {
	GetEpisodeById(id int) (*Episode, error)
}
type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

type RegisterUserPayLoad struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=12"`
}
type LoginUserPayLoad struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type Movie struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `json:"name"`
	Year        int    `json:"year"`
	NumEpisodes int    `json:"numEpisodes"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Country     string `json:"country"`
	Thumbnail   string `json:"thumbnail"`
	Trailer     string `json:"trailer"`
	isFree      bool   `json:"isFree"`
}
type Episode struct {
	ID             int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Episode_Number int    `gorm:"not null;uniqueIndex:idx_episode_movie" json:"episode_number"`
	MovieID        int    `gorm:"not null;uniqueIndex:idx_episode_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_id"`
	Source         string `json:"source"`
	Duration       int    `json:"duration"`
	Quality        string `json:"quality"`
}

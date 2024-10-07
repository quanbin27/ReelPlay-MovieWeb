package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int64) (*User, error)
	CreateUser(user *User) error
}
type MovieStore interface {
	GetMoviesWithPagination(offset, limit int) ([]MovieResponse, error)
}
type EpisodeStore interface {
	GetEpisodeById(id int) (*Episode, error)
}
type RateStore interface{}
type FavoriteStore interface{}
type CommentStore interface{}
type CategoryStore interface{}
type ActorStore interface{}
type DirectorStore interface{}
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
	ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `json:"name"`
	Year        int        `json:"year"`
	NumEpisodes int        `json:"numEpisodes"`
	Description string     `json:"description"`
	Language    string     `json:"language"`
	Country     string     `json:"country"`
	Thumbnail   string     `json:"thumbnail"`
	Trailer     string     `json:"trailer"`
	Rate        float32    `json:"rate"`
	IsFree      bool       `json:"is_free"`
	Category    []Category `json:"category" gorm:"many2many:movie_category"`
	Actor       []Actor    `json:"actor" gorm:"many2many:movie_actor"`
	Director    []Director `json:"director" gorm:"many2many:movie_director"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdateAt    time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
}
type MovieResponse struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Year        int      `json:"year"`
	NumEpisodes int      `json:"numEpisodes"`
	Description string   `json:"description"`
	Language    string   `json:"language"`
	Country     string   `json:"country"`
	Thumbnail   string   `json:"thumbnail"`
	Trailer     string   `json:"trailer"`
	Rate        float32  `json:"rate"`
	IsFree      bool     `json:"is_free"`
	Category    []string `json:"category"`
	Actor       []string `json:"actor"`
	Director    []string `json:"director"`
}
type Director struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Actor struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
}
type Category struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
}
type Episode struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	EpisodeNumber int       `gorm:"not null;uniqueIndex:idx_episode_movie" json:"episode_number"`
	MovieID       int       `gorm:"not null;uniqueIndex:idx_episode_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_id"`
	Source        string    `json:"source"`
	Duration      int       `json:"duration"`
	Quality       string    `json:"quality"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdateAt      time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
type Comment struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string    `gorm:"not null" json:"content"`
	UserID    int       `gorm:"not null;uniqueIndex:idx_user_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	MovieID   int       `gorm:"not null;uniqueIndex:idx_user_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}
type Favourite struct {
	UserID    int       `gorm:"not null;primaryKey:idx_user_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	MovieID   int       `gorm:"not null;primaryKey:idx_user_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}
type Rate struct {
	UserID    int       `gorm:"not null;primaryKey:idx_user_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	MovieID   int       `gorm:"not null;primaryKey:idx_user_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_id"`
	Rate      int       `gorm:"not null" json:"rate"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

package types

import "time"

type WatchStore interface {
	UpdateWatchPosition(userId, episodeId, position int) error
	GetWatchPosition(userId, episodeId int) (int, error)
	CreateUserWatched(userId int, episodeId int) error
}
type CategoryFitStore interface {
	GetUserCategoryFit(userID int, categoryID int) (*CategoryFit, error)
	CreateCategoryFit(fit *CategoryFit) error
	UpdateCategoryFit(fit *CategoryFit) error
}
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
	UpdatePassword(userID int, newPassword string) error
}
type EmailService interface {
	SendResetPasswordEmail(to, resetLink string) error
}
type MovieStore interface {
	GetMoviesWithPagination(offset, limit int) ([]MovieResponse, error)
	GetMovieById(id int) (MovieResponse, error)
	GetCategories(id int) ([]int, error)
}
type EpisodeStore interface {
	GetEpisodeByMovieAndEpisodeId(movieId, episodeNumber int) (*Episode, error)
	GetEpisodeById(id int) (*Episode, error)
}
type RateStore interface {
	CreateMovieRating(rating *Rate) error
	UpdateMovieRating(rating *Rate) error
	GetUserMovieRating(userID int, movieID int) (*Rate, error)
	UpdateMovieAverageRating(movieID int) error
}
type BookmarkStore interface {
	GetBookmarksByUser(userId int) ([]Bookmark, error)
	IsBookmark(movieId int, userId int) bool
	CreateBookmark(movieId int, userId int) error
	CancelBookmark(movieId int, userId int) error
}
type CommentStore interface {
	CreateComment(content string, movieID int, userID int) (*Comment, error)
	GetCommentsByMovieID(movieID int) ([]Comment, error)
}
type CategoryStore interface{}
type ActorStore interface{}
type DirectorStore interface{}

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

type CreateCommentRequest struct {
	Content string `json:"content" `
	MovieID int    `json:"movie_id" `
	UserID  int    `json:"user_id" `
}
type CommentResponse struct {
	ID        int    `json:"id"`
	MovieID   int    `json:"movie_id"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"` // This will hold the user's name
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type CreateBookmarkRequest struct {
	MovieID int `json:"movie_id" `
	UserID  int `json:"user_id" `
}

// User struct
type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// Movie struct
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
	Category    []Category `gorm:"many2many:movie_category"`
	Actor       []Actor    `gorm:"many2many:movie_actor"`
	Director    []Director `gorm:"many2many:movie_director"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

// Episode struct
type Episode struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	EpisodeNumber int       `gorm:"not null;uniqueIndex:idx_episode_movie" json:"episode_number"`
	MovieID       int       `gorm:"not null;uniqueIndex:idx_episode_movie;foreignKey:MovieID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_id"`
	Source        string    `json:"source"`
	Duration      int       `json:"duration"`
	Quality       string    `json:"quality"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Movie         Movie     `gorm:"foreignKey:MovieID"`
}

// Bookmark struct
type Bookmark struct {
	UserID    int       `gorm:"not null;primaryKey:idx_user_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	MovieID   int       `gorm:"not null;primaryKey:idx_user_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Movie     Movie     `gorm:"foreignKey:MovieID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// Comment struct
type Comment struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string    `gorm:"not null" json:"content"`
	UserID    int       `gorm:"not null;foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	MovieID   int       `gorm:"not null;foreignKey:MovieID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	User      User      `gorm:"foreignKey:UserID"`
	Movie     Movie     `gorm:"foreignKey:MovieID"`
}

// Rate struct
type Rate struct {
	UserID    int       `gorm:"not null;primaryKey:idx_user_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	MovieID   int       `gorm:"not null;primaryKey:idx_user_movie;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_id"`
	Rate      int       `gorm:"not null" json:"rate"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Movie     Movie     `gorm:"foreignKey:MovieID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// Category Fit Struct
type CategoryFit struct {
	UserID     int       `gorm:"not null;primaryKey:idx_user_category;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	CategoryID int       `gorm:"not null;primaryKey:idx_user_category;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category_id"`
	FitRate    float32   `gorm:"not null" json:"fit_rate"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	User       User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Category   Category  `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// Director struct
type Director struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
}

// Actor struct
type Actor struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
}

// Category struct
type Category struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
}

// UserWatched struct
type UserWatched struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int       `gorm:"not null;foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	EpisodeID    int       `gorm:"not null;foreignKey:EpisodeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"episode_id"`
	LastPosition int       `gorm:"not null;default:0" json:"last_position"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	User         User      `gorm:"foreignKey:UserID"`
	Episode      Episode   `gorm:"foreignKey:EpisodeID"`
}

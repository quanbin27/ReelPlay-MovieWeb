package types

import (
	"gorm.io/gorm"
	"time"
)

type WatchStore interface {
	UpdateWatchPosition(userId, episodeId, position int) error
	GetWatchPosition(userId, episodeId int) (int, error)
	CreateUserWatched(userId int, episodeId int) error
}
type CategoryFitStore interface {
	GetUserCategoryFit(userID int, categoryID int) (*CategoryFit, error)
	CreateCategoryFit(fit *CategoryFit) error
	UpdateCategoryFit(fit *CategoryFit) error
	GetCategoryFit(userId int) (cate1, cate2, cate3 *CategoryFit, err error)
}
type DirectorStore interface {
	GetAllDirectors() ([]Director, error)
	DeleteDirector(id int) error
	UpdateDirector(id int, updatedDirector *Director) error
	GetDirectorByID(id int) (*Director, error)
	CreateDirector(director *Director) error
	SearchDirectors(keyword string, page, limit int) ([]Director, int64, error)
}
type ActorStore interface {
	GetAllActors() ([]Actor, error)
	DeleteActor(id int) error
	UpdateActor(id int, updatedActor *Actor) error
	GetActorByID(id int) (*Actor, error)
	CreateActor(director *Actor) error
	SearchActors(keyword string, page, limit int) ([]Actor, int64, error)
}
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
	UpdatePassword(userID int, newPassword string) error
	SearchUsers(keyword string, page, limit int) ([]User, int64, error)
	DeleteUserSoft(userID int) error
	UpdateUserInfo(userID int, updatedData map[string]interface{}) error
	UpdateUserPassword(userID int, newPassword string) error
	UnlockUser(userID int) error
	CountUsers() (int, error)
}
type EmailService interface {
	SendResetPasswordEmail(to, resetLink string) error
}
type MovieStore interface {
	GetNewMovies(limit int) ([]MovieItemResponse, error)
	GetMostViewRates(limit int) ([]MovieItemResponse, error)
	GetMostViewMovies(limit int) ([]MovieItemResponse, error)
	IncrementView(movieID int) error
	GetAllMovies() ([]Movie, error)
	GetMoviesWithPagination(offset, limit int) ([]MovieItemResponse, error)
	GetMovieById(id int) (MovieResponse, error)
	GetCategories(id int) ([]int, error)
	MovieSearch(query string, offset, limit int) ([]MovieItemResponse, error)
	MovieSearchCount(query string) (int64, error)
	GetMoviesByCategories(userId, cate1Id, cate2Id, cate3Id int) ([]MovieItemResponse, error)
	GetNewRecommendedMovies(userId, cate1Id, cate2Id, cate3Id int) ([]MovieItemResponse, error)
	CreateMovie(movie *Movie, categoryIDs []int, actorIDs []int, directorIDs []int) error
	UpdateMovie(id int, updateReq *UpdateMovieRequest) error
	DeleteMovie(id int) error
	UpdateNumofEp(movieId int, num int) error
	UpdateAverageDuration(movieId int) error
	CountMovies() (int, error)
	CountViews() (int, error)
	SumRates() (int, error)
}
type EpisodeStore interface {
	GetEpisodeByMovieAndEpisodeId(movieId, episodeNumber int) (*Episode, error)
	GetEpisodeById(id int) (*Episode, error)
	DeleteEpisode(id int) error
	UpdateEpisode(id int, episode *UpdateEpisodeRequest) error
	CreateEpisode(episode *Episode) error
	SearchEpisodes(keyword string, page int, limit int) ([]SearchEpisodesResponse, int64, error)
	CountEpisodes() (int, error)
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
	GetCommentsByUserID(userID int) ([]Comment, error)
	DeleteComment(commentID int) error
}
type CategoryStore interface{}

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
	View        int      `json:"view"`
	IsRecommend bool     `json:"is_recommend"`
	Category    []string `json:"category"`
	Actor       []string `json:"actor"`
	Director    []string `json:"director"`
}
type MovieItemResponse struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Rate      float32  `json:"rate"`
	Category  []string `json:"category"`
	View      int      `json:"view"`
	Thumbnail string   `json:"thumbnail"`
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
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	RoleID    int            `json:"role_id"` // Không cần dùng foreignKey ở đây
	Role      Role           `gorm:"foreignKey:RoleID" json:"role"`
}

type Role struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"unique" json:"name"`
}

// Movie struct
type Movie struct {
	ID            int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string         `json:"name"`
	Year          int            `json:"year"`
	NumEpisodes   int            `json:"num_episodes"`
	Description   string         `json:"description"`
	Language      string         `json:"language"`
	CountryID     int            `json:"country_id"`
	TimeForEp     int            `json:"time_for_ep"`
	Thumbnail     string         `json:"thumbnail"`
	Trailer       string         `json:"trailer"`
	Rate          float32        `json:"rate"`
	PredictRate   float32        `json:"predict_rate"`
	View          int            `json:"view" gorm:"default:0" `
	IsRecommended bool           `json:"is_recommended"`
	Category      []Category     `gorm:"many2many:movie_category"`
	Actor         []Actor        `gorm:"many2many:movie_actor"`
	Director      []Director     `gorm:"many2many:movie_director"`
	Country       Country        `gorm:"foreignKey:CountryID"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
type CreateMovieInput struct {
	Name          string  `json:"name" validate:"required"`
	Year          int     `json:"year" validate:"required"`
	Description   string  `json:"description"`
	Language      string  `json:"language"`
	CountryID     int     `json:"country_id" validate:"required"`
	Thumbnail     string  `json:"thumbnail"`
	Trailer       string  `json:"trailer"`
	PredictRate   float32 `json:"predict_rate"`
	IsRecommended bool    `json:"is_recommended"`
	CategoryIDs   []int   `json:"category_ids" validate:"required"`
	ActorIDs      []int   `json:"actor_ids" validate:"required"`
	DirectorIDs   []int   `json:"director_ids" validate:"required"`
}
type AllMovieResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type UpdateMovieRequest struct {
	Name          string `json:"name"`
	Year          int    `json:"year"`
	Description   string `json:"description"`
	Language      string `json:"language"`
	CountryID     int    `json:"country_id"`
	Thumbnail     string `json:"thumbnail"`
	Trailer       string `json:"trailer"`
	IsRecommended bool   `json:"is_recommended"`
	CategoryIds   []int  `json:"category_ids"` // Sử dụng CategoryIds
	ActorIds      []int  `json:"actor_ids"`    // Sử dụng ActorIds
	DirectorIds   []int  `json:"director_ids"` // Sử dụng DirectorIds
}
type Country struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

// Episode struct
type Episode struct {
	ID            int            `gorm:"primaryKey;autoIncrement" json:"id"`
	EpisodeNumber int            `gorm:"not null;uniqueIndex:idx_episode_movie" json:"episode_number"`
	MovieID       int            `gorm:"not null;uniqueIndex:idx_episode_movie" json:"movie_id"`
	Source        string         `json:"source"`
	Duration      int            `json:"duration"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Movie         Movie          `gorm:"foreignKey:MovieID"`
}
type UpdateEpisodeRequest struct {
	Source   string `json:"source"`
	Duration int    `json:"duration"`
}
type SearchEpisodesResponse struct {
	ID            int       `json:"id"`
	MovieID       int       `json:"movie_id"`
	EpisodeNumber int       `json:"episode_number"`
	MovieTitle    string    `json:"movie_title"`
	Duration      int       `json:"duration"`
	UpdatedAt     time.Time `json:"updated_at"`
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
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string         `gorm:"not null" json:"content"`
	UserID    int            `gorm:"not null;foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	MovieID   int            `gorm:"not null;foreignKey:MovieID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	User      User           `gorm:"foreignKey:UserID"`
	Movie     Movie          `gorm:"foreignKey:MovieID"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
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
	Year int    `json:"year"`
}

// Actor struct
type Actor struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
	Year int    `json:"year"`
}

type Category struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserWatched struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int       `gorm:"not null;foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
	EpisodeID    int       `gorm:"not null;foreignKey:EpisodeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"episode_id"`
	LastPosition int       `gorm:"not null;default:0" json:"last_position"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	User         User      `gorm:"foreignKey:UserID"`
	Episode      Episode   `gorm:"foreignKey:EpisodeID"`
}

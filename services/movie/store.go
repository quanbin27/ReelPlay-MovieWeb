package movie

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
func (s *Store) GetMoviesWithPagination(offset, limit int) ([]types.MovieItemResponse, error) {
	var movies []types.Movie
	result := s.db.Preload("Category").Limit(limit).Offset(offset).Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range movies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate

		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}

		// Thêm movieResponse vào slice kết quả
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}
func (s *Store) GetMoviesByCategories(userId, cate1Id, cate2Id, cate3Id int) ([]types.MovieItemResponse, error) {
	var movies []types.Movie

	// Lấy danh sách movieId đã xem dựa trên episodeId từ bảng user_watcheds
	var watchedMovieIds []int
	s.db.Table("user_watcheds").
		Joins("LEFT JOIN episodes ON user_watcheds.episode_id = episodes.id").
		Where("user_watcheds.user_id = ?", userId).
		Pluck("episodes.movie_id", &watchedMovieIds)
	print("da xem: ", watchedMovieIds[0], "!")
	// Khởi tạo truy vấn cơ bản
	query := s.db.Preload("Category").Joins("LEFT JOIN movie_category ON movies.id = movie_category.movie_id")

	// Thêm điều kiện is_recommended
	query = query.Where("movies.is_recommended = ?", 1)

	// Thêm điều kiện loại bỏ phim đã xem nếu danh sách watchedMovieIds không rỗng
	if len(watchedMovieIds) > 0 {
		query = query.Where("movies.id NOT IN ?", watchedMovieIds)
	}

	// Trường hợp không có thể loại nào, lấy tất cả các phim is_recommended
	if cate1Id == 0 && cate2Id == 0 && cate3Id == 0 {
		query = query.Group("movies.id").Order("movies.rate DESC").Find(&movies)
	} else {
		// Trường hợp có 1, 2 hoặc 3 category
		condition := ""
		if cate2Id == 0 { // Chỉ có 1 category
			condition = "movie_category.category_id = ?"
			query = query.Where(condition, cate1Id)
		} else if cate3Id == 0 { // Có 2 category
			condition = "(movie_category.category_id = ? OR movie_category.category_id = ?)"
			query = query.Where(condition, cate1Id, cate2Id)
		} else { // Có 3 category
			condition = "(movie_category.category_id = ? OR movie_category.category_id = ? OR movie_category.category_id = ?)"
			query = query.Where(condition, cate1Id, cate2Id, cate3Id)
		}

		query = query.Group("movies.id").Order(`
			CASE 
				WHEN COUNT(DISTINCT movie_category.category_id) = 3 THEN 1
				WHEN COUNT(DISTINCT movie_category.category_id) = 2 THEN 2
				ELSE 3 
			END, movies.rate DESC`).Find(&movies)
	}

	// Kiểm tra lỗi từ kết quả truy vấn
	if query.Error != nil {
		return nil, query.Error
	}

	// Chuyển đổi kết quả phim sang dạng MovieItemResponse
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range movies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate
		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}

func (s *Store) GetNewRecommendedMovies(userId, cate1Id, cate2Id, cate3Id int) ([]types.MovieItemResponse, error) {
	var movies []types.Movie

	var watchedMovieIds []int
	s.db.Table("user_watcheds").
		Joins("LEFT JOIN episodes ON user_watcheds.episode_id = episodes.id").
		Where("user_watcheds.user_id = ?", userId).
		Pluck("episodes.movie_id", &watchedMovieIds)
	print(len(watchedMovieIds), " ", watchedMovieIds[0])
	for id := range watchedMovieIds {
		print("movie id: ", id)
	}
	// Xác định khoảng thời gian phim mới được tạo (trong vòng 10 ngày)
	tenDaysAgo := time.Now().AddDate(0, 0, -10)

	// Trường hợp không có thể loại nào, lấy tất cả các phim mới, is_recommended, và sắp xếp theo predict_rate
	if cate1Id == 0 && cate2Id == 0 && cate3Id == 0 {
		result := s.db.Preload("Category").
			Joins("LEFT JOIN movie_category ON movies.id = movie_category.movie_id").
			Where("movies.is_recommended = ? AND movies.created_at >= ?", 1, tenDaysAgo).
			Where("movies.id NOT IN ?", watchedMovieIds). // Loại bỏ phim đã xem
			Group("movies.id").
			Order("movies.predict_rate DESC").
			Find(&movies)
		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		// Trường hợp có 1, 2 hoặc 3 category
		condition := ""
		if cate2Id == 0 { // Chỉ có 1 category
			condition = "movie_category.category_id = ?"
		} else if cate3Id == 0 { // Có 2 category
			condition = "(movie_category.category_id = ? OR movie_category.category_id = ?)"
		} else { // Có 3 category
			condition = "(movie_category.category_id = ? OR movie_category.category_id = ? OR movie_category.category_id = ?)"
		}

		result := s.db.Preload("Category").
			Joins("LEFT JOIN movie_category ON movies.id = movie_category.movie_id").
			Where(condition, cate1Id, cate2Id, cate3Id).
			Where("movies.is_recommended = ? AND movies.created_at >= ?", 1, tenDaysAgo).
			Where("movies.id NOT IN ?", watchedMovieIds). // Loại bỏ phim đã xem
			Group("movies.id").
			Order(`
				CASE 
					WHEN COUNT(DISTINCT movie_category.category_id) = 3 THEN 1
					WHEN COUNT(DISTINCT movie_category.category_id) = 2 THEN 2
					ELSE 3 
				END, movies.predict_rate DESC
			`).
			Find(&movies)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	// Chuyển đổi kết quả phim sang dạng MovieItemResponse
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range movies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate
		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}

		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}

func (s *Store) GetCategories(id int) ([]int, error) {
	var movie types.Movie
	result := s.db.Preload("Category").Where("id = ?", id).First(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	var categories []int
	for _, category := range movie.Category {
		categories = append(categories, category.ID)
	}
	return categories, nil
}
func (s *Store) GetMovieById(id int) (types.MovieResponse, error) {
	var movie types.Movie
	result := s.db.Preload("Actor").Preload("Category").Preload("Director").Preload("Country").Where("id = ?", id).First(&movie)
	if result.Error != nil {
		return types.MovieResponse{}, result.Error
	}
	var movieResponse types.MovieResponse
	movieResponse.ID = movie.ID
	movieResponse.Name = movie.Name
	movieResponse.Year = movie.Year
	movieResponse.NumEpisodes = movie.NumEpisodes
	movieResponse.Description = movie.Description
	movieResponse.Language = movie.Language
	movieResponse.Country = movie.Country.Name
	movieResponse.Thumbnail = movie.Thumbnail
	movieResponse.Trailer = movie.Trailer
	movieResponse.Rate = movie.Rate

	for _, actor := range movie.Actor {
		movieResponse.Actor = append(movieResponse.Actor, actor.Name)
	}

	for _, director := range movie.Director {
		movieResponse.Director = append(movieResponse.Director, director.Name)
	}

	for _, category := range movie.Category {
		movieResponse.Category = append(movieResponse.Category, category.Name)
	}
	return movieResponse, nil
}
func (s *Store) MovieSearchCount(query string) (int64, error) {
	var count int64
	result := s.db.Table("movies").
		Joins("LEFT JOIN movie_actor ON movies.id = movie_actor.movie_id").
		Joins("LEFT JOIN actors ON movie_actor.actor_id = actors.id").
		Joins("LEFT JOIN movie_director ON movies.id = movie_director.movie_id").
		Joins("LEFT JOIN directors ON movie_director.director_id = directors.id").
		Where("movies.name LIKE ? OR actors.name LIKE ? OR directors.name LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Group("movies.id").Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (s *Store) MovieSearch(query string, offset, limit int) ([]types.MovieItemResponse, error) {
	var movies []types.Movie
	result := s.db.Table("movies").Preload("Category").
		Joins("LEFT JOIN movie_actor ON movies.id = movie_actor.movie_id").
		Joins("LEFT JOIN actors ON movie_actor.actor_id = actors.id").
		Joins("LEFT JOIN movie_director ON movies.id = movie_director.movie_id").
		Joins("LEFT JOIN directors ON movie_director.director_id = directors.id").
		Where("movies.name LIKE ? OR actors.name LIKE ? OR directors.name LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Group("movies.id").Limit(limit).Offset(offset).
		Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range movies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate
		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}

		// Thêm movieResponse vào slice kết quả
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}

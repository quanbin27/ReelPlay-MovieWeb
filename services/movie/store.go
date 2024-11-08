package movie

import (
	"fmt"
	"github.com/quanbin27/ReelPlay/types"
	"gorm.io/gorm"
	"strconv"
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
		movieItemResponse.View = movie.View
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
		movieItemResponse.View = movie.View

		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}

// UpdateMovie cập nhật thông tin movie theo ID
func (s *Store) UpdateMovie(id int, updateReq *types.UpdateMovieRequest) error {
	// Kiểm tra xem movie có tồn tại không
	var existingMovie types.Movie
	if err := s.db.First(&existingMovie, id).Error; err != nil {
		return err // Nếu không tìm thấy movie, trả về lỗi
	}

	// Cập nhật thông tin movie
	existingMovie.Name = updateReq.Name
	existingMovie.Year = updateReq.Year
	existingMovie.Description = updateReq.Description
	existingMovie.Language = updateReq.Language
	existingMovie.CountryID = updateReq.CountryID
	existingMovie.Thumbnail = updateReq.Thumbnail
	existingMovie.Trailer = updateReq.Trailer
	existingMovie.IsRecommended = updateReq.IsRecommended

	// Cập nhật movie trong cơ sở dữ liệu
	if err := s.db.Save(&existingMovie).Error; err != nil {
		return err // Nếu có lỗi khi cập nhật, trả về lỗi
	}

	// Cập nhật các mối quan hệ
	// Cập nhật Diễn viên (Actor)
	if len(updateReq.ActorIds) > 0 {
		var actors []types.Actor
		if err := s.db.Find(&actors, updateReq.ActorIds).Error; err != nil {
			return err
		}
		if err := s.db.Model(&existingMovie).Association("Actor").Replace(actors); err != nil {
			return err
		}
	} else {
		// Xóa tất cả liên kết diễn viên nếu không có ActorIds
		if err := s.db.Model(&existingMovie).Association("Actor").Clear(); err != nil {
			return err
		}
	}

	// Cập nhật Đạo diễn (Director)
	if len(updateReq.DirectorIds) > 0 {
		var directors []types.Director
		if err := s.db.Find(&directors, updateReq.DirectorIds).Error; err != nil {
			return err
		}
		if err := s.db.Model(&existingMovie).Association("Director").Replace(directors); err != nil {
			return err
		}
	} else {
		// Xóa tất cả liên kết đạo diễn nếu không có DirectorIds
		if err := s.db.Model(&existingMovie).Association("Director").Clear(); err != nil {
			return err
		}
	}

	// Cập nhật Thể loại (Category)
	if len(updateReq.CategoryIds) > 0 {
		var categories []types.Category
		if err := s.db.Find(&categories, updateReq.CategoryIds).Error; err != nil {
			return err
		}
		if err := s.db.Model(&existingMovie).Association("Category").Replace(categories); err != nil {
			return err
		}
	} else {
		// Xóa tất cả liên kết thể loại nếu không có CategoryIds
		if err := s.db.Model(&existingMovie).Association("Category").Clear(); err != nil {
			return err
		}
	}

	return nil // Cập nhật thành công
}

func (s *Store) DeleteMovie(id int) error {
	// Kiểm tra xem movie có tồn tại không
	var movie types.Movie
	if err := s.db.First(&movie, id).Error; err != nil {
		return err // Nếu không tìm thấy movie, trả về lỗi
	}

	// Thực hiện xóa mềm
	if err := s.db.Delete(&movie).Error; err != nil {
		return err // Nếu có lỗi khi xóa, trả về lỗi
	}

	return nil // Xóa thành công
}
func (s *Store) UpdateNumofEp(movieId int, num int) error {
	var movie types.Movie
	if err := s.db.First(&movie, movieId).Error; err != nil {
		return err // Nếu không tìm thấy movie, trả về lỗi
	}
	newNum := movie.NumEpisodes + num
	if err := s.db.Model(&types.Movie{}).Where("id = ?", movieId).Update("num_episodes", newNum).Error; err != nil {
		return err // Nếu có lỗi khi xóa, trả về lỗi
	}
	return nil
}
func (s *Store) GetAllMovies() ([]types.Movie, error) {
	var movies []types.Movie
	// Chỉ chọn cột ID và Name
	err := s.db.Select("id", "name").Find(&movies).Error
	if err != nil {
		return nil, err
	}
	return movies, nil
}
func (s *Store) UpdateAverageDuration(movieId int) error {
	// Tính tổng thời gian của tất cả các tập phim
	var totalDuration int
	var episodeCount int

	// Lấy tất cả các tập phim của movieId
	var episodes []types.Episode
	if err := s.db.Where("movie_id = ? AND deleted_at IS NULL", movieId).Find(&episodes).Error; err != nil {
		return err
	}

	// Tính tổng thời gian và số lượng tập
	for _, episode := range episodes {
		totalDuration += episode.Duration
		episodeCount++
	}

	// Tính thời gian trung bình
	var averageDuration int
	if episodeCount > 0 {
		averageDuration = totalDuration / episodeCount
	} else {
		averageDuration = 0 // Nếu không có tập nào, đặt trung bình là 0
	}

	// Cập nhật vào bảng movies
	if err := s.db.Model(&types.Movie{}).Where("id = ?", movieId).Update("time_for_ep", averageDuration).Error; err != nil {
		return err
	}

	return nil
}

func (s *Store) CreateMovie(movie *types.Movie, categoryIDs []int, actorIDs []int, directorIDs []int) error {
	// Bắt đầu một transaction để đảm bảo tính toàn vẹn của dữ liệu
	tx := s.db.Begin()

	// Tạo movie trong database
	if err := tx.Create(&movie).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Thiết lập các quan hệ nhiều-nhiều (categories)
	if len(categoryIDs) > 0 {
		var categories []types.Category
		if err := tx.Where("id IN ?", categoryIDs).Find(&categories).Error; err != nil {
			tx.Rollback()
			return err
		}
		// GORM sẽ tự động tạo bản ghi vào bảng trung gian (movie_category)
		if err := tx.Model(&movie).Association("Category").Append(&categories); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Thiết lập các quan hệ nhiều-nhiều (actors)
	if len(actorIDs) > 0 {
		var actors []types.Actor
		if err := tx.Where("id IN ?", actorIDs).Find(&actors).Error; err != nil {
			tx.Rollback()
			return err
		}
		// GORM sẽ tự động tạo bản ghi vào bảng trung gian (movie_actor)
		if err := tx.Model(&movie).Association("Actor").Append(&actors); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Thiết lập các quan hệ nhiều-nhiều (directors)
	if len(directorIDs) > 0 {
		var directors []types.Director
		if err := tx.Where("id IN ?", directorIDs).Find(&directors).Error; err != nil {
			tx.Rollback()
			return err
		}
		// GORM sẽ tự động tạo bản ghi vào bảng trung gian (movie_director)
		if err := tx.Model(&movie).Association("Director").Append(&directors); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction nếu mọi thứ đều thành công
	return tx.Commit().Error
}
func (s *Store) GetMostViewMovies(limit int) ([]types.MovieItemResponse, error) {
	var movies []types.Movie
	if err := s.db.Preload("Category").Order("view DESC").Limit(limit).Find(&movies).Error; err != nil {
		return nil, err
	}
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range movies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate
		movieItemResponse.View = movie.View

		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}
func (s *Store) GetNewMovies(limit int) ([]types.MovieItemResponse, error) {
	var movies []types.Movie
	if err := s.db.Preload("Category").Order("created_at DESC").Limit(limit).Find(&movies).Error; err != nil {
		return nil, err
	}
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range movies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate
		movieItemResponse.View = movie.View

		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}
func (s *Store) GetFeaturesMovies(limit int) ([]types.MovieItemResponse, error) {
	var movies []types.Movie
	if err := s.db.Preload("Category").Where("num_episodes = ?", 1).Order("created_at DESC").Limit(limit).Find(&movies).Error; err != nil {
		return nil, err
	}
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range movies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate
		movieItemResponse.View = movie.View

		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}
func (s *Store) GetSeriesMovies(limit int) ([]types.MovieItemResponse, error) {
	var movies []types.Movie
	if err := s.db.Preload("Category").Where("num_episodes > ?", 1).Order("created_at DESC").Limit(limit).Find(&movies).Error; err != nil {
		return nil, err
	}
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range movies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate
		movieItemResponse.View = movie.View

		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}
func (s *Store) GetBookMarkMovies(userID int) ([]types.MovieItemResponse, error) {
	var bookmarkedMovies []types.Movie

	// Lấy danh sách các bộ phim đã bookmark của người dùng
	err := s.db.Joins("JOIN bookmarks ON bookmarks.movie_id = movies.id").
		Where("bookmarks.user_id = ?", userID).
		Preload("Category").
		Find(&bookmarkedMovies).Error

	if err != nil {
		return nil, err
	}

	// Chuyển đổi từ []Movie sang []MovieItemResponse
	var result []types.MovieItemResponse
	for _, movie := range bookmarkedMovies {
		// Lấy tên các danh mục
		var categories []string
		for _, category := range movie.Category {
			categories = append(categories, category.Name)
		}

		// Tạo MovieItemResponse
		result = append(result, types.MovieItemResponse{
			ID:        movie.ID,
			Name:      movie.Name,
			Rate:      movie.Rate,
			Category:  categories,
			View:      movie.View,
			Thumbnail: movie.Thumbnail,
		})
	}

	return result, nil
}

func (s *Store) CountMovies() (int, error) {
	var count int64                                                  // Thử dùng int64 thay cho int
	if err := s.db.Table("movies").Count(&count).Error; err != nil { // Chỉ định tên bảng trực tiếp
		return 0, err
	}
	return int(count), nil
}
func (s *Store) SumRates() (int, error) {
	var sum int64
	if err := s.db.Model(&types.Rate{}).Select("COALESCE(SUM(rate), 0)").Scan(&sum).Error; err != nil {
		return 0, err
	}
	return int(sum), nil
}

func (s *Store) CountViews() (int, error) {
	var totalViews int64

	// Tính tổng views từ tất cả các phim trong bảng movie
	if err := s.db.Table("movies").Select("COALESCE(SUM(view), 0)").Scan(&totalViews).Error; err != nil {
		return 0, fmt.Errorf("failed to count total movie views: %w", err)
	}

	return int(totalViews), nil
}

func (s *Store) GetMostViewRates(limit int) ([]types.MovieItemResponse, error) {
	var movies []types.Movie
	if err := s.db.Preload("Category").Order("rate DESC").Limit(limit).Find(&movies).Error; err != nil {
		return nil, err
	}
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range movies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate
		movieItemResponse.View = movie.View

		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}
func (s *Store) GetWatchingList(userId int) ([]types.MovieItemResponse, error) {
	var watchedMovies []types.Movie

	err := s.db.Unscoped().Table("user_watcheds").
		Joins("LEFT JOIN episodes ON user_watcheds.episode_id = episodes.id").
		Joins("LEFT JOIN movies ON episodes.movie_id = movies.id").
		Where("user_watcheds.user_id = ?", userId).
		Where("movies.deleted_at IS NULL").
		Preload("Category").
		Select("DISTINCT movies.*").
		Find(&watchedMovies).Error

	if err != nil {
		return nil, err
	}
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range watchedMovies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate
		movieItemResponse.View = movie.View

		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil

}
func (s *Store) GetNewRecommendedMovies(userId, cate1Id, cate2Id, cate3Id int) ([]types.MovieItemResponse, error) {
	var movies []types.Movie

	// Lấy danh sách movieId đã xem dựa trên episodeId từ bảng user_watcheds
	var watchedMovieIds []int
	s.db.Table("user_watcheds").
		Joins("LEFT JOIN episodes ON user_watcheds.episode_id = episodes.id").
		Where("user_watcheds.user_id = ?", userId).
		Pluck("episodes.movie_id", &watchedMovieIds)

	// Khởi tạo truy vấn cơ bản
	query := s.db.Preload("Category").
		Joins("LEFT JOIN movie_category ON movies.id = movie_category.movie_id")

	// Thêm điều kiện là phim được recommended
	query = query.Where("movies.is_recommended = ?", 1)

	// Xác định khoảng thời gian phim mới được tạo (trong vòng 10 ngày)
	tenDaysAgo := time.Now().AddDate(0, 0, -10)
	query = query.Where("movies.created_at >= ?", tenDaysAgo)

	// Thêm điều kiện loại bỏ phim đã xem nếu danh sách watchedMovieIds không rỗng
	if len(watchedMovieIds) > 0 {
		query = query.Where("movies.id NOT IN ?", watchedMovieIds)
	}

	// Trường hợp không có thể loại nào, lấy tất cả các phim mới
	if cate1Id == 0 && cate2Id == 0 && cate3Id == 0 {
		query = query.Group("movies.id").
			Order("movies.predict_rate DESC").
			Find(&movies)
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

		// Sắp xếp phim theo số lượng category khớp và predict_rate
		query = query.Group("movies.id").
			Order(`
				CASE 
					WHEN COUNT(DISTINCT movie_category.category_id) = 3 THEN 1
					WHEN COUNT(DISTINCT movie_category.category_id) = 2 THEN 2
					ELSE 3 
				END, movies.predict_rate DESC`).
			Find(&movies)
	}

	// Kiểm tra lỗi từ kết quả truy vấn
	if query.Error != nil {
		return nil, query.Error
	}
	if len(movies) < 5 {
		var additionalMovies []types.Movie
		// Truy vấn bổ sung để lấy các phim mới không bị xem, không cần điều kiện theo category
		s.db.Preload("Category").
			Where("movies.is_recommended = ?", 1).
			Where("movies.created_at >= ?", time.Now().AddDate(0, 0, -30)).
			Not("movies.id", watchedMovieIds).
			Order("movies.predict_rate DESC").
			Limit(5 - len(movies)). // Lấy số lượng phim cần thêm để đạt ít nhất 5
			Find(&additionalMovies)

		// Thêm các phim bổ sung vào danh sách
		movies = append(movies, additionalMovies...)
	}

	// Chuyển đổi kết quả phim sang dạng MovieItemResponse
	var movieItemResponses []types.MovieItemResponse
	for _, movie := range movies {
		var movieItemResponse types.MovieItemResponse
		movieItemResponse.ID = movie.ID
		movieItemResponse.Name = movie.Name
		movieItemResponse.Thumbnail = movie.Thumbnail
		movieItemResponse.Rate = movie.Rate
		movieItemResponse.View = movie.View

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
func (s *Store) IncrementView(movieID int) error {
	return s.db.Model(&types.Movie{}).Where("id = ?", movieID).Update("view", gorm.Expr("view + ?", 1)).Error
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
	movieResponse.IsRecommend = movie.IsRecommended
	movieResponse.View = movie.View

	for _, actor := range movie.Actor {
		movieResponse.Actor = append(movieResponse.Actor, actor.Name+" "+strconv.Itoa(actor.Year))
	}

	for _, director := range movie.Director {
		movieResponse.Director = append(movieResponse.Director, director.Name+" "+strconv.Itoa(director.Year))
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
		Joins("LEFT JOIN movie_category ON movies.id = movie_category.movie_id").
		Joins("LEFT JOIN categories ON movie_category.category_id = categories.id").
		Joins("LEFT JOIN countries ON movies.country_id = countries.id").
		Where("movies.deleted_at IS NULL").
		Where("movies.name LIKE ? OR actors.name LIKE ? OR directors.name LIKE ? OR categories.name LIKE ? OR countries.name LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").
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
		Joins("LEFT JOIN movie_category ON movies.id = movie_category.movie_id").
		Joins("LEFT JOIN categories ON movie_category.category_id = categories.id").
		Joins("LEFT JOIN countries ON movies.country_id = countries.id").
		Where("movies.name LIKE ? OR actors.name LIKE ? OR directors.name LIKE ? OR categories.name LIKE ? OR countries.name LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").
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
		movieItemResponse.View = movie.View
		for _, category := range movie.Category {
			movieItemResponse.Category = append(movieItemResponse.Category, category.Name)
		}

		// Thêm movieResponse vào slice kết quả
		movieItemResponses = append(movieItemResponses, movieItemResponse)
	}

	return movieItemResponses, nil
}

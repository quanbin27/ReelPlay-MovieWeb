package movie

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/services/auth"
	"github.com/quanbin27/ReelPlay/types"
	"github.com/quanbin27/ReelPlay/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Handler struct {
	UserStore        types.UserStore
	MovieStore       types.MovieStore
	CategoryFitStore types.CategoryFitStore
}

func NewHandler(userStore types.UserStore, moviestore types.MovieStore, categoryFitStore types.CategoryFitStore) *Handler {
	return &Handler{userStore, moviestore, categoryFitStore}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {

	//	e.GET("/movie", h.GetMovies)
	e.GET("/movie/:id", h.GetMovieByID)
	e.GET("/movie/:id/category", h.GetCategoryID)
	e.GET("/movie/search", h.MovieSearch)
	e.GET("/movie/most-views/:limit", h.GetMostViewMovies)
	e.GET("/movie/most-rates/:limit", h.GetMostViewRates)
	e.GET("/movie/new/:limit", h.NewMovie)
	e.GET("/movie/new/series/:limit", h.GetSeriesMovies)
	e.GET("/movie/new/features/:limit", h.GetFeaturesMovies)
	e.GET("/movies", h.GetAllMovies)
	e.GET("/movie/user/recommend", h.GetRecommendedMoviesByCategory, auth.WithJWTAuth(h.UserStore))
	e.GET("/movie/user/new-recommend", h.GetNewRecommendedMovies, auth.WithJWTAuth(h.UserStore))
	e.GET("/movie/new-user/recommend", h.GetRecommendedMoviesByCategoryNew)
	e.GET("/movie/new-user/new-recommend", h.GetNewRecommendedMoviesNew)
	e.PUT("/movie/:id", h.UpdateMovieHandler, auth.WithJWTAdminAuth(h.UserStore))
	e.GET("/dashboard", h.GetDashBoardIndex, auth.WithJWTAdminAuth(h.UserStore))
	e.POST("/movie", h.CreateMovie, auth.WithJWTAdminAuth(h.UserStore))
	e.DELETE("/movie/:id", h.DeleteMovieHandler, auth.WithJWTAdminAuth(h.UserStore))
	e.GET("/me/bookmarks", h.GetBookMarkList, auth.WithJWTAuth(h.UserStore))
	e.GET("/me/watching", h.getWatching, auth.WithJWTAuth(h.UserStore))
}
func (h *Handler) getWatching(c echo.Context) error {
	userId, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Can't get userId from context"})
	}
	watchingList, err := h.MovieStore.GetWatchingList(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Can't get watching list"})
	}
	return c.JSON(http.StatusOK, watchingList)
}
func (h *Handler) UpdateMovieHandler(c echo.Context) error {
	// Lấy ID từ path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	// Tạo đối tượng request để lưu thông tin mới
	var updateReq types.UpdateMovieRequest
	if err := c.Bind(&updateReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	if err := h.MovieStore.UpdateMovie(id, &updateReq); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "movie not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update movie"})
	}

	// Trả về phản hồi thành công
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "movie updated successfully",
	})
}

func (h *Handler) CreateMovie(c echo.Context) error {
	var input types.CreateMovieInput

	// Bind dữ liệu từ request vào struct CreateMovieInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	// Validate dữ liệu đầu vào
	if err := utils.Validate.Struct(&input); err != nil {
		errors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": errors[0].Error()})
	}

	// Tạo đối tượng Movie để lưu vào database
	movie := types.Movie{
		Name:          input.Name,
		Year:          input.Year,
		Description:   input.Description,
		Language:      input.Language,
		CountryID:     input.CountryID,
		Thumbnail:     input.Thumbnail,
		Trailer:       input.Trailer,
		PredictRate:   input.PredictRate,
		IsRecommended: input.IsRecommended,
	}

	// Lưu movie vào store
	err := h.MovieStore.CreateMovie(&movie, input.CategoryIDs, input.ActorIDs, input.DirectorIDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not save movie",
		})
	}

	return c.JSON(http.StatusCreated, movie)
}
func (h *Handler) DeleteMovieHandler(c echo.Context) error {
	// Lấy ID từ path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	// Gọi hàm xóa trong store
	if err := h.MovieStore.DeleteMovie(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "movie not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete movie"})
	}

	// Trả về phản hồi thành công
	return c.NoContent(http.StatusNoContent) // 204 No Content
}
func (h *Handler) GetRecommendedMoviesByCategory(c echo.Context) error {
	userId, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}

	cate1, cate2, cate3, err := h.CategoryFitStore.GetCategoryFit(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching user category fit"})
	}

	var cate1Id, cate2Id, cate3Id int
	if cate1 != nil {
		cate1Id = cate1.CategoryID
	}
	if cate2 != nil {
		cate2Id = cate2.CategoryID
	}
	if cate3 != nil {
		cate3Id = cate3.CategoryID
	}
	movies, err := h.MovieStore.GetMoviesByCategories(userId, cate1Id, cate2Id, cate3Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching movies"})
	}

	return c.JSON(http.StatusOK, movies)
}
func (h *Handler) GetRecommendedMoviesByCategoryNew(c echo.Context) error {
	movies, err := h.MovieStore.GetMoviesByCategories(0, 0, 0, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching movies"})
	}
	return c.JSON(http.StatusOK, movies)
}
func (h *Handler) GetBookMarkList(c echo.Context) error {
	userId, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't get user Id"})
	}
	bookmarkMovieList, err := h.MovieStore.GetBookMarkMovies(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't get bookmarks list"})
	}
	return c.JSON(http.StatusOK, bookmarkMovieList)
}

// Route cho lấy danh sách các phim mới được đề xuất
func (h *Handler) GetNewRecommendedMovies(c echo.Context) error {
	userId, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}

	// Lấy thông tin category phù hợp của user
	cate1, cate2, cate3, err := h.CategoryFitStore.GetCategoryFit(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching user category fit"})
	}

	var cate1Id, cate2Id, cate3Id int
	if cate1 != nil {
		cate1Id = cate1.CategoryID
	}
	if cate2 != nil {
		cate2Id = cate2.CategoryID
	}
	if cate3 != nil {
		cate3Id = cate3.CategoryID
	}
	print(cate1Id, cate2Id, cate3Id)
	// Gọi store để lấy danh sách phim mới được đề xuất
	movies, err := h.MovieStore.GetNewRecommendedMovies(userId, cate1Id, cate2Id, cate3Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching new recommended movies"})
	}

	return c.JSON(http.StatusOK, movies)
}
func (h *Handler) GetNewRecommendedMoviesNew(c echo.Context) error {

	movies, err := h.MovieStore.GetNewRecommendedMovies(0, 0, 0, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching new recommended movies"})
	}

	return c.JSON(http.StatusOK, movies)
}

func (h *Handler) GetCategoryID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	var categories []int
	categories, err = h.MovieStore.GetCategories(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, categories)
}
func (h *Handler) GetMovieByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	movieResponse, err := h.MovieStore.GetMovieById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, movieResponse)
}
func (h *Handler) NewMovie(c echo.Context) error {
	limitParam := c.Param("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 18
	}
	movies, err := h.MovieStore.GetNewMovies(limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not get new movies"})
	}
	return c.JSON(http.StatusOK, movies)
}
func (h *Handler) MovieSearch(c echo.Context) error {
	query := c.QueryParam("q")
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	totalResults, err := h.MovieStore.MovieSearchCount(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not count movies"})
	}
	movies, err := h.MovieStore.MovieSearch(query, offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not search movies"})
	}
	response := map[string]interface{}{
		"movies": movies,
		"total":  totalResults,
		"page":   page,
		"limit":  limit,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetMovies(c echo.Context) error {

	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	movies, err := h.MovieStore.GetMoviesWithPagination(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch movies",
		})
	}
	return c.JSON(http.StatusOK, movies)
}
func (h *Handler) GetAllMovies(c echo.Context) error {
	// Lấy danh sách phim từ store
	movies, err := h.MovieStore.GetAllMovies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch movies"})
	}

	// Chuyển đổi sang định dạng MovieResponse để trả về
	var movieResponses []types.AllMovieResponse
	for _, movie := range movies {
		movieResponses = append(movieResponses, types.AllMovieResponse{
			ID:   movie.ID,
			Name: movie.Name,
		})
	}

	return c.JSON(http.StatusOK, movieResponses)
}
func (h *Handler) GetDashBoardIndex(c echo.Context) error {
	views, err := h.MovieStore.CountViews()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	movies, err := h.MovieStore.CountMovies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	rates, err := h.MovieStore.SumRates()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	user, err := h.UserStore.CountUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"views":  views,
		"movies": movies,
		"rates":  rates,
		"user":   user,
	})
}
func (h *Handler) GetMostViewMovies(c echo.Context) error {
	limit, _ := strconv.Atoi(c.Param("limit"))
	movies, err := h.MovieStore.GetMostViewMovies(limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, movies)
}
func (h *Handler) GetMostViewRates(c echo.Context) error {
	limit, _ := strconv.Atoi(c.Param("limit"))
	movies, err := h.MovieStore.GetMostViewRates(limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, movies)
}
func (h *Handler) GetFeaturesMovies(c echo.Context) error {
	limit, _ := strconv.Atoi(c.Param("limit"))
	movies, err := h.MovieStore.GetFeaturesMovies(limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, movies)
}
func (h *Handler) GetSeriesMovies(c echo.Context) error {
	limit, _ := strconv.Atoi(c.Param("limit"))
	movies, err := h.MovieStore.GetSeriesMovies(limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, movies)
}

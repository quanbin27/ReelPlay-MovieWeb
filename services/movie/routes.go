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
	CategoryStore    types.CategoryStore
	ActorStore       types.ActorStore
	DirectorStore    types.DirectorStore
	CategoryFitStore types.CategoryFitStore
}

func NewHandler(userStore types.UserStore, moviestore types.MovieStore, categorystore types.CategoryStore, actorstore types.ActorStore, directorstore types.DirectorStore, categoryFitStore types.CategoryFitStore) *Handler {
	return &Handler{userStore, moviestore, categorystore, actorstore, directorstore, categoryFitStore}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.GET("/movie", h.GetAllMovies)
	//	e.GET("/movie", h.GetMovies)
	e.GET("/movie/:id", h.GetMovieByID)
	e.GET("/movie/:id/category", h.GetCategoryID)
	e.GET("/movie/search", h.MovieSearch)
	e.GET("/movie/user/:user_id/recommend", h.GetRecommendedMoviesByCategory, auth.WithJWTAuth(h.UserStore))
	e.GET("/movie/user/:user_id/new-recommend", h.GetNewRecommendedMovies, auth.WithJWTAuth(h.UserStore))
	e.POST("/movie", h.CreateMovie, auth.WithJWTAdminAuth(h.UserStore))
	e.DELETE("/movie/:id", h.DeleteMovieHandler, auth.WithJWTAdminAuth(h.UserStore))
	e.PUT("/movie/:id", h.UpdateMovieHandler, auth.WithJWTAdminAuth(h.UserStore))
}

// UpdateMovie cập nhật thông tin movie theo ID
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

	// Gọi hàm cập nhật trong store
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
	userId, err := strconv.Atoi(c.Param("user_id"))
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

// Route cho lấy danh sách các phim mới được đề xuất
func (h *Handler) GetNewRecommendedMovies(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
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
	if query == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Query parameter is required"})
	}
	totalResults, err := h.MovieStore.MovieSearchCount(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not count movies"})
	}
	movies, err := h.MovieStore.MovieSearch(query, offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not search movies"})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": movies, "totalResults": totalResults})
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

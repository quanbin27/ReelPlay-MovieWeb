package movie

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/types"
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
	e.GET("/movies", h.GetMovies)
	e.GET("/movie/:id", h.GetMovieByID)
	e.GET("/movie/:id/category", h.GetCategoryID)
	e.GET("/movie/search", h.MovieSearch)
	e.GET("/movie/user/:user_id/recommend", h.GetRecommendedMoviesByCategory)
	e.GET("/movie/user/:user_id/new-recommend", h.GetNewRecommendedMovies)

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

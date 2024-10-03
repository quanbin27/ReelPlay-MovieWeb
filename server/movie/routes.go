package movie

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
	"strconv"
)

type Handler struct {
	store types.MovieStore
}

func NewHandler(store types.MovieStore) *Handler {
	return &Handler{store}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.GET("/movies", h.GetMovies)
	e.GET("/movie", h.GetMoviestest)
}
func (h *Handler) GetMoviestest(c echo.Context) error {
	return c.String(http.StatusOK, "Movies")
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

	movies, err := h.store.GetMoviesWithPagination(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch movies",
		})
	}

	return c.JSON(http.StatusOK, movies)
}

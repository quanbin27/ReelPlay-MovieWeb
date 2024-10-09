package movie

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
	"strconv"
)

type Handler struct {
	MovieStore    types.MovieStore
	CategoryStore types.CategoryStore
	ActorStore    types.ActorStore
	DirectorStore types.DirectorStore
}

func NewHandler(moviestore types.MovieStore, categorystore types.CategoryStore, actorstore types.ActorStore, directorstore types.DirectorStore) *Handler {
	return &Handler{moviestore, categorystore, actorstore, directorstore}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.GET("/movies", h.GetMovies)
	e.GET("/movie/:id", h.GetMovieByID)
}
func (h *Handler) GetMovieByID(c echo.Context) error {
	id := c.Param("id")
	movieResponse, err := h.MovieStore.GetMovieById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, movieResponse)
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
	//var moviesresponse []types.MovieResponse
	//for i := 0; i < len(movies); i++ {
	//	moviesresponse[i].ID = movies[i].ID
	//	moviesresponse[i].Name = movies[i].Name
	//	moviesresponse[i].Year = movies[i].Year
	//	moviesresponse[i].NumEpisodes = movies[i].NumEpisodes
	//	moviesresponse[i].Description = movies[i].Description
	//	moviesresponse[i].Language = movies[i].Language
	//	moviesresponse[i].Country = movies[i].Country
	//	moviesresponse[i].Thumbnail = movies[i].Thumbnail
	//	moviesresponse[i].Trailer = movies[i].Trailer
	//	moviesresponse[i].Rate = movies[i].Rate
	//	moviesresponse[i].IsFree = movies[i].IsFree
	//
	//}
	return c.JSON(http.StatusOK, movies)
}

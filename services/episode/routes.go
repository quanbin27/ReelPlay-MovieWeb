package episode

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/services/auth"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
	"strconv"
)

type Handler struct {
	episodeStore types.EpisodeStore
	userStore    types.UserStore
	movieStore   types.MovieStore
}

func NewHandler(episodeStore types.EpisodeStore, userStore types.UserStore, movieStore types.MovieStore) *Handler {
	return &Handler{episodeStore, userStore, movieStore}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	g := e.Group("", auth.WithJWTAuth(h.userStore))
	//g.GET("/episode/:id", h.GetEpisodeById)
	g.GET("/movie/:movieId/episode/:episodeNumber", h.GetEpisodeByMovieAndEpisodeId)

}

func (h *Handler) GetEpisodeByMovieAndEpisodeId(c echo.Context) error {
	movieIdStr := c.Param("movieId")
	episodeNumberStr := c.Param("episodeNumber")

	movieId, err := strconv.Atoi(movieIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid movie ID"})
	}

	episodeNumber, err := strconv.Atoi(episodeNumberStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid episode number"})
	}

	episode, err := h.episodeStore.GetEpisodeByMovieAndEpisodeId(movieId, episodeNumber)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, episode)
}

//func (h *Handler) GetEpisodeById(c echo.Context) error {
//	idStr := c.Param("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid episode ID"})
//	}
//	episode, err := h.episodeStore.GetEpisodeById(id)
//	if err != nil {
//		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
//	}
//	return c.JSON(http.StatusOK, episode)
//}

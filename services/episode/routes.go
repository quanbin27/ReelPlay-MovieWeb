package episode

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/services/auth"
	"github.com/quanbin27/ReelPlay/types"
	"gorm.io/gorm"
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

	e.GET("/episode/:id", h.GetEpisodeById, auth.WithJWTAuth(h.userStore))
	e.GET("/movie/:movieId/episode/:episodeNumber", h.GetEpisodeByMovieAndEpisodeId, auth.WithJWTAuth(h.userStore))
	e.POST("/episode", h.CreateEpisodeHandler, auth.WithJWTAuth(h.userStore))
	e.PUT("/episode/:id", h.UpdateEpisodeHandler, auth.WithJWTAuth(h.userStore))
	e.DELETE("/episode/:id", h.DeleteEpisodeHandler, auth.WithJWTAuth(h.userStore))

}
func (h *Handler) CreateEpisodeHandler(c echo.Context) error {
	var episode types.Episode
	if err := c.Bind(&episode); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	if err := h.episodeStore.CreateEpisode(&episode); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create episode"})
	}

	return c.JSON(http.StatusCreated, episode)
}
func (h *Handler) UpdateEpisodeHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var episode types.Episode
	if err := c.Bind(&episode); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	if err := h.episodeStore.UpdateEpisode(id, &episode); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "episode not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update episode"})
	}

	return c.JSON(http.StatusOK, episode)
}

func (h *Handler) DeleteEpisodeHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	if err := h.episodeStore.DeleteEpisode(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "episode not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete episode"})
	}

	return c.NoContent(http.StatusNoContent) // Trả về No Content
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

func (h *Handler) GetEpisodeById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid episode ID"})
	}
	episode, err := h.episodeStore.GetEpisodeById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, episode)
}

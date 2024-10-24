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
	e.POST("/episode", h.CreateEpisodeHandler, auth.WithJWTAdminAuth(h.userStore))
	e.PUT("/episode/:id", h.UpdateEpisodeHandler, auth.WithJWTAdminAuth(h.userStore))
	e.DELETE("/episode/:id", h.DeleteEpisodeHandler, auth.WithJWTAdminAuth(h.userStore))
	e.GET("/episode", h.SearchEpisodesHandler, auth.WithJWTAdminAuth(h.userStore))

}
func (h *Handler) CreateEpisodeHandler(c echo.Context) error {
	var episode types.Episode
	if err := c.Bind(&episode); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	if err := h.episodeStore.CreateEpisode(&episode); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create episode"})
	}
	if err := h.movieStore.UpdateNumofEp(episode.MovieID, 1); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't update Num of Episode"})
	}
	if err := h.movieStore.UpdateAverageDuration(episode.MovieID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't Update Average Duration"})
	}

	return c.JSON(http.StatusCreated, episode)
}
func (h *Handler) SearchEpisodesHandler(c echo.Context) error {
	// Get query params for pagination and search
	keyword := c.QueryParam("keyword")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Fetch episodes with pagination and search
	episodes, total, err := h.episodeStore.SearchEpisodes(keyword, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch episodes",
		})
	}

	// Create pagination metadata
	response := map[string]interface{}{
		"episodes": episodes,
		"page":     page,
		"limit":    limit,
		"total":    total,
	}

	// Return response
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateEpisodeHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	episode, err := h.episodeStore.GetEpisodeById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update episode"})
	}
	var newEpisode types.UpdateEpisodeRequest
	if err := c.Bind(&newEpisode); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	if err := h.episodeStore.UpdateEpisode(id, &newEpisode); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "episode not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update episode"})
	}

	if err := h.movieStore.UpdateAverageDuration(episode.MovieID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't Update Average Duration"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "episode updated"})
}

func (h *Handler) DeleteEpisodeHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	episode, err := h.episodeStore.GetEpisodeById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "can't find episode"})
	}

	if err := h.episodeStore.DeleteEpisode(episode.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete episode"})
	}
	if err := h.movieStore.UpdateNumofEp(episode.MovieID, -1); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't update Num of Episode"})
	}
	if err := h.movieStore.UpdateAverageDuration(episode.MovieID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't Update Average Duration"})
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

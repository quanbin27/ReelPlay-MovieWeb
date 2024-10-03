package episode

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/server/auth"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
	"strconv"
)

type Handler struct {
	episodeStore types.EpisodeStore
	userStore    types.UserStore
}

func NewHandler(episodeStore types.EpisodeStore, userStore types.UserStore) *Handler {
	return &Handler{episodeStore, userStore}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	g := e.Group("", auth.WithJWTAuth(h.userStore))
	g.GET("/episode/:id", h.GetEpisodeById)
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

package user_watched

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/services/auth"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
	"strconv"
)

type Handler struct {
	userStore    types.UserStore
	watchStore   types.WatchStore
	episodeStore types.EpisodeStore
	movieStore   types.MovieStore
}

func NewHandler(store types.UserStore, watchStore types.WatchStore, episodeStore types.EpisodeStore, movieStore types.MovieStore) *Handler {
	return &Handler{store, watchStore, episodeStore, movieStore}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	g := e.Group("", auth.WithJWTAuth(h.userStore))
	g.PUT("/watch-position", h.UpdateWatchPosition)
	g.GET("/watch-position", h.GetWatchPosition)
	g.POST("/watch-position", h.CreateUserWatched)
}
func (h *Handler) UpdateWatchPosition(c echo.Context) error {
	userId, _ := strconv.Atoi(c.QueryParam("user_id"))
	episodeId, _ := strconv.Atoi(c.QueryParam("episode_id"))
	lastPosition, _ := strconv.Atoi(c.QueryParam("position"))

	err := h.watchStore.UpdateWatchPosition(userId, episodeId, lastPosition)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	print("Da update time xem", lastPosition)
	return c.JSON(http.StatusOK, echo.Map{"message": "Watch position updated"})
}

func (h *Handler) GetWatchPosition(c echo.Context) error {
	userIdStr := c.QueryParam("user_id")
	episodeIdStr := c.QueryParam("episode_id")

	// Chuyển đổi các giá trị từ chuỗi sang số nguyên và kiểm tra lỗi
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user_id"})
	}
	episodeId, err := strconv.Atoi(episodeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid episode_id"})
	}

	// Gọi đến store để lấy vị trí xem
	position, err := h.watchStore.GetWatchPosition(userId, episodeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"last_position": position})
}
func (h *Handler) CreateUserWatched(c echo.Context) error {
	userIdStr := c.QueryParam("user_id")
	episodeIdStr := c.QueryParam("episode_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user_id"})
	}
	episodeId, err := strconv.Atoi(episodeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid episode_id"})
	}

	// Lấy Episode để truy xuất MovieID

	episode, err := h.episodeStore.GetEpisodeById(episodeId)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Episode not found"})
	}
	// Tăng view cho Movie
	if err := h.movieStore.IncrementView(episode.MovieID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Unable to increment movie view"})
	}

	// Tạo bản ghi UserWatched
	err = h.watchStore.CreateUserWatched(userId, episodeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   "Unable to create user watched record",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User watched record created"})
}

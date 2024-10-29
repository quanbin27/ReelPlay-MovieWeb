package bookmark

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/services/auth"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
)

type Handler struct {
	bookmarkStore types.BookmarkStore
	userStore     types.UserStore
}

func NewHandler(store types.BookmarkStore, userStore types.UserStore) *Handler {
	return &Handler{store, userStore}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.POST("/movie/bookmark", h.Create, auth.WithJWTAuth(h.userStore))
	e.DELETE("/movie/bookmark", h.CancelBookmark, auth.WithJWTAuth(h.userStore))
	e.POST("/bookmark/exist", h.IsBookmark, auth.WithJWTAuth(h.userStore))

}

func (h *Handler) Create(c echo.Context) error {
	req := new(types.CreateBookmarkRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	err := h.bookmarkStore.CreateBookmark(req.MovieID, req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create bookmark"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Bookmark created successfully"})
}
func (h *Handler) CancelBookmark(c echo.Context) error {
	req := new(types.CreateBookmarkRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	err := h.bookmarkStore.CancelBookmark(req.MovieID, req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to cancel bookmark"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Bookmark cancelled"})
}
func (h *Handler) IsBookmark(c echo.Context) error {
	req := new(types.CreateBookmarkRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	result := h.bookmarkStore.IsBookmark(req.MovieID, req.UserID)
	return c.JSON(http.StatusOK, result)
}

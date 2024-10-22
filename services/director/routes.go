package director

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
	directorStore types.DirectorStore
	userStore     types.UserStore
}

func NewHandler(directorStore types.DirectorStore, userStore types.UserStore) *Handler {
	return &Handler{directorStore, userStore}
}

// DirectorRoutes defines all routes for director-related operations
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.POST("/director", h.CreateDirectorHandler, auth.WithJWTAdminAuth(h.userStore))

	e.PUT("/director/:id", h.UpdateDirectorHandler, auth.WithJWTAdminAuth(h.userStore))

	e.DELETE("/director/:id", h.DeleteDirectorHandler, auth.WithJWTAdminAuth(h.userStore))

	e.GET("/director", h.GetAllDirectorsHandler)
	e.GET("/director/:id", h.GetDirectorByIDHandler)

}
func (h *Handler) CreateDirectorHandler(c echo.Context) error {
	// Tạo đối tượng Director mới
	var director types.Director
	if err := c.Bind(&director); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	// Gọi hàm trong store để tạo Director
	if err := h.directorStore.CreateDirector(&director); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create director"})
	}

	// Trả về phản hồi thành công
	return c.JSON(http.StatusCreated, director)
}
func (h *Handler) UpdateDirectorHandler(c echo.Context) error {
	// Lấy ID từ path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	// Tạo đối tượng Director để cập nhật thông tin
	var updatedDirector types.Director
	if err := c.Bind(&updatedDirector); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	// Gọi hàm trong store để cập nhật Director
	if err := h.directorStore.UpdateDirector(id, &updatedDirector); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "director not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update director"})
	}

	// Trả về phản hồi thành công
	return c.JSON(http.StatusOK, updatedDirector)
}
func (h *Handler) DeleteDirectorHandler(c echo.Context) error {
	// Lấy ID từ path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	// Gọi hàm trong store để xóa Director
	if err := h.directorStore.DeleteDirector(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "director not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete director"})
	}

	// Trả về phản hồi thành công
	return c.NoContent(http.StatusNoContent)
}
func (h *Handler) GetAllDirectorsHandler(c echo.Context) error {
	// Gọi hàm trong store để lấy tất cả Directors
	directors, err := h.directorStore.GetAllDirectors()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch directors"})
	}

	// Trả về danh sách Directors
	return c.JSON(http.StatusOK, directors)
}
func (h *Handler) GetDirectorByIDHandler(c echo.Context) error {
	// Lấy ID từ path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	// Gọi store để lấy thông tin của Director
	director, err := h.directorStore.GetDirectorByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "director not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch director"})
	}

	// Trả về thông tin của Director
	return c.JSON(http.StatusOK, director)
}

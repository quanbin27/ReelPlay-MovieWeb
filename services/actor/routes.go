package actor

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
	ActorStore types.ActorStore
	userStore  types.UserStore
}

func NewHandler(ActorStore types.ActorStore, userStore types.UserStore) *Handler {
	return &Handler{ActorStore, userStore}
}

// ActorRoutes defines all routes for Actor-related operations
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.POST("/actor", h.CreateActorHandler, auth.WithJWTAdminAuth(h.userStore))

	e.PUT("/actor/:id", h.UpdateActorHandler, auth.WithJWTAdminAuth(h.userStore))

	e.DELETE("/actor/:id", h.DeleteActorHandler, auth.WithJWTAdminAuth(h.userStore))

	e.GET("/actors", h.GetAllActorsHandler)
	e.GET("/actor/:id", h.GetActorByIDHandler)
	e.GET("/actor", h.SearchActorsHandler)
}
func (h *Handler) SearchActorsHandler(c echo.Context) error {
	// Lấy keyword từ query params
	keyword := c.QueryParam("keyword")

	// Lấy các giá trị phân trang từ query params
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1 // Mặc định là trang 1 nếu không có giá trị
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10 // Mặc định là 10 bản ghi mỗi trang
	}

	// Gọi store để lấy danh sách diễn viên và tổng số bản ghi
	actors, total, err := h.ActorStore.SearchActors(keyword, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch actors"})
	}

	// Trả về dữ liệu JSON bao gồm danh sách diễn viên và tổng số bản ghi
	response := map[string]interface{}{
		"actors": actors,
		"total":  total,
		"page":   page,
		"limit":  limit,
	}
	return c.JSON(http.StatusOK, response)
}
func (h *Handler) CreateActorHandler(c echo.Context) error {
	// Tạo đối tượng Actor mới
	var Actor types.Actor
	if err := c.Bind(&Actor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	// Gọi hàm trong store để tạo Actor
	if err := h.ActorStore.CreateActor(&Actor); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create Actor"})
	}

	// Trả về phản hồi thành công
	return c.JSON(http.StatusCreated, Actor)
}
func (h *Handler) UpdateActorHandler(c echo.Context) error {
	// Lấy ID từ path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	// Tạo đối tượng Actor để cập nhật thông tin
	var updatedActor types.Actor
	if err := c.Bind(&updatedActor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	// Gọi hàm trong store để cập nhật Actor
	if err := h.ActorStore.UpdateActor(id, &updatedActor); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Actor not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update Actor"})
	}

	// Trả về phản hồi thành công
	return c.JSON(http.StatusOK, updatedActor)
}
func (h *Handler) DeleteActorHandler(c echo.Context) error {
	// Lấy ID từ path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	// Gọi hàm trong store để xóa Actor
	if err := h.ActorStore.DeleteActor(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Actor not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete Actor"})
	}

	// Trả về phản hồi thành công
	return c.NoContent(http.StatusNoContent)
}
func (h *Handler) GetAllActorsHandler(c echo.Context) error {
	// Gọi hàm trong store để lấy tất cả Actors
	Actors, err := h.ActorStore.GetAllActors()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch Actors"})
	}

	// Trả về danh sách Actors
	return c.JSON(http.StatusOK, Actors)
}
func (h *Handler) GetActorByIDHandler(c echo.Context) error {
	// Lấy ID từ path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	// Gọi store để lấy thông tin của Actor
	Actor, err := h.ActorStore.GetActorByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Actor not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch Actor"})
	}

	// Trả về thông tin của Actor
	return c.JSON(http.StatusOK, Actor)
}

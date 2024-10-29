package category_fit

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/services/auth"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
	"strconv"
)

type Handler struct {
	fitStore  types.CategoryFitStore
	userStore types.UserStore
}

func NewHandler(fitStore types.CategoryFitStore, userStore types.UserStore) *Handler {
	return &Handler{fitStore, userStore}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.POST("/user/category-fit", h.CategoryFit, auth.WithJWTAuth(h.userStore))
}
func (h *Handler) CategoryFit(c echo.Context) error {
	userId, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID"})
	}

	fitValue, err := strconv.Atoi(c.FormValue("fit_rate"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid fit rate"})
	}

	categoryId, err := strconv.Atoi(c.FormValue("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid category ID"})
	}

	existingFit, err := h.fitStore.GetUserCategoryFit(userId, categoryId)
	if err != nil && err.Error() != "record not found" {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not check existing category rating"})
	}

	if existingFit != nil {
		// Cập nhật đánh giá hiện tại
		existingFit.FitRate = (existingFit.FitRate + float32(fitValue)) / 2
		if err := h.fitStore.UpdateCategoryFit(existingFit); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not update category rating"})
		}
	} else {
		// Tạo đánh giá mới
		newRating := types.CategoryFit{
			UserID:     userId,
			CategoryID: categoryId,
			FitRate:    float32(fitValue),
		}
		if err := h.fitStore.CreateCategoryFit(&newRating); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not create category rating"})
		}
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Category fit successfully created"})
}

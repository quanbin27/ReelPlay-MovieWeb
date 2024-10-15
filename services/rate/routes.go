package rate

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/services/auth"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
	"strconv"
)

type Handler struct {
	rateStore types.RateStore
	userStore types.UserStore
}

func NewHandler(store types.RateStore, userStore types.UserStore) *Handler {
	return &Handler{store, userStore}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {

	e.POST("/movie/:movieId/rate", h.RateMovie, auth.WithJWTAuth(h.userStore)) // Create or update movie rating
}
func (h *Handler) RateMovie(c echo.Context) error {
	print("vao day")
	userId, _ := strconv.Atoi(c.QueryParam("user_id"))
	movieId, _ := strconv.Atoi(c.Param("movieId"))
	ratingValue, _ := strconv.Atoi(c.FormValue("rate"))
	print(userId, movieId, ratingValue)
	// Check if user has already rated the movie
	existingRating, err := h.rateStore.GetUserMovieRating(userId, movieId)
	if err != nil && err.Error() != "record not found" {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not check existing rating"})
	}

	// If the rating exists, update it
	if existingRating != nil {
		existingRating.Rate = ratingValue
		if err := h.rateStore.UpdateMovieRating(existingRating); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not update rating"})
		}
	} else {
		// Otherwise, create a new rating
		newRating := types.Rate{
			UserID:  userId,
			MovieID: movieId,
			Rate:    ratingValue,
		}
		if err := h.rateStore.CreateMovieRating(&newRating); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not create rating"})
		}
	}

	// Tính lại rating trung bình cho phim
	if err := h.rateStore.UpdateMovieAverageRating(movieId); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not update movie rating"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Rating saved and movie rating updated"})
}

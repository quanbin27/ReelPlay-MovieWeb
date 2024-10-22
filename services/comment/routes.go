package comment

import (
	"github.com/labstack/echo/v4"
	"github.com/quanbin27/ReelPlay/services/auth"
	"github.com/quanbin27/ReelPlay/types"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	cmtStore  types.CommentStore
	userStore types.UserStore
}

func NewHandler(cmtStore types.CommentStore, userStore types.UserStore) *Handler {
	return &Handler{cmtStore, userStore}
}
func (h *Handler) RegisterRoutes(e *echo.Group) {
	e.POST("/comment", h.Create, auth.WithJWTAuth(h.userStore))
	e.GET("/movie/:id/comment", h.GetCommentsByMovieID)
}
func (h *Handler) Create(c echo.Context) error {
	req := new(types.CreateCommentRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	comment, err := h.cmtStore.CreateComment(req.Content, req.MovieID, req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create comment"})
	}
	user, err := h.userStore.GetUserByID(req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create comment"})
	}
	cmtreponse := types.CommentResponse{
		ID:        comment.ID,
		MovieID:   comment.MovieID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt.Format(time.RFC3339),
		Content:   comment.Content,
		UserName:  user.FirstName + " " + user.LastName,
	}

	return c.JSON(http.StatusCreated, cmtreponse)
}
func (h *Handler) GetCommentsByMovieID(c echo.Context) error {
	movieIDParam := c.Param("id")
	movieID, err := strconv.Atoi(movieIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid movie ID"})
	}

	comments, err := h.cmtStore.GetCommentsByMovieID(movieID)
	var commentResponses []types.CommentResponse

	for _, comment := range comments {
		user, err := h.userStore.GetUserByID(comment.UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch user for comment",
			})
		}
		commentResponses = append(commentResponses, types.CommentResponse{
			ID:        comment.ID,
			MovieID:   comment.MovieID,
			UserID:    comment.UserID,
			UserName:  user.FirstName + " " + user.LastName,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Format(time.RFC3339),
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch comments"})
	}

	return c.JSON(http.StatusOK, commentResponses)
}

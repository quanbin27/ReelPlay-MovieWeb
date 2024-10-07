package comment

import "github.com/quanbin27/ReelPlay/types"

type Handler struct {
	store types.CommentStore
}

func NewHandler(store types.CommentStore) *Handler {
	return &Handler{store}
}

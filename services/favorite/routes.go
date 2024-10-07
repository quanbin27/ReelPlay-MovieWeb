package favorite

import "github.com/quanbin27/ReelPlay/types"

type Handler struct {
	store types.FavoriteStore
}

func NewHandler(store types.FavoriteStore) *Handler {
	return &Handler{store}
}

package rate

import "github.com/quanbin27/ReelPlay/types"

type Handler struct {
	store types.RateStore
}

func NewHandler(store types.RateStore) *Handler {
	return &Handler{store}
}

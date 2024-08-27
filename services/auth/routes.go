package auth

import "github.com/gorilla/mux"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) AuthRoutes(router *mux.Router) {

}

package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	UserHandler *UserHandler
}

func NewHandler(uh *UserHandler) *Handler {
	return &Handler{UserHandler: uh}
}

func (h *Handler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/users/{id}", h.UserHandler.HandleGetUser)
	return r
}

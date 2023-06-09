package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/andreashanson/go-template/internal/user"
	"github.com/go-chi/chi/v5"
)

type UserService interface {
	GetUser(ctx context.Context, id string) (user.User, error)
}

type UserHandler struct {
	us UserService
}

func NewUserHandler(us UserService) *UserHandler {
	return &UserHandler{us: us}
}

func (uh *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Add("Content-Type", "application/json")

	u, err := uh.us.GetUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

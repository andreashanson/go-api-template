package mongo

import (
	"context"

	"github.com/andreashanson/go-template/internal/user"
)

type UserRepo struct {
	Conn      *Connection
	FakeStore Store
}

func NewUserRepo(c *Connection) *UserRepo {
	s := Store{users: []user.User{{ID: "1", Name: "Andreas"}, {ID: "2", Name: "Anna"}}}
	return &UserRepo{Conn: c, FakeStore: s}
}

func (r *UserRepo) GetUserByID(ctx context.Context, id string) (user.User, error) {
	for _, u := range r.FakeStore.users {
		if u.ID == id {
			return u, nil
		}
	}

	return user.User{}, user.ErrUserNotFound
}

type Store struct {
	users []user.User
}

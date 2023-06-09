package user

import (
	"context"
)

type Repository interface {
	GetUserByID(ctx context.Context, id string) (User, error)
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetUser(ctx context.Context, id string) (User, error) {
	return s.repo.GetUserByID(ctx, id)
}

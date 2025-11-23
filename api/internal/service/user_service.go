package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/secret-santa-rubot/api/internal/domain"
	"github.com/secret-santa-rubot/api/internal/model"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, name, username string) error {
	user := &model.User{
		UserId:   uuid.New().String(),
		Name:     name,
		Username: username,
	}

	return s.repo.Create(ctx, user)
}

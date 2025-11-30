package service

import (
	"context"
	"fmt"
	"github.com/secret-santa-rubot/api/internal/domain"
	"github.com/secret-santa-rubot/api/internal/model"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, userId, name, username string) (string, error) {
	user := &model.User{
		UserId:   userId,
		Name:     name,
		Username: username,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}
	return userId, nil
}

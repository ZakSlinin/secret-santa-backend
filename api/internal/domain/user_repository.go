package domain

import (
	"context"
	"github.com/secret-santa-rubot/api/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, u *model.User) error
}

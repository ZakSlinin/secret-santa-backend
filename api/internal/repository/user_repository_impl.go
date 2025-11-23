package repository

import (
	"context"
	"database/sql"
	"github.com/secret-santa-rubot/api/internal/model"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(ctx context.Context, u *model.User) error {
	query := `INSERT INTO users (user_id, name, username, groups_member) VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, u.UserId, u.Name, u.Username, u.GroupsMember)
	return err
}

package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/LevinStrike/lux-backend/internal/core/entities"
)

func (r *Repository) CreateUser(ctx context.Context, username, password string) (entities.User, error) {
	row := r.db.QueryRowContext(ctx, "INSERT INTO users(email, password) VALUES($1, $2) RETURNING id;", username, password)
	var id int
	err := row.Scan(&id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entities.User{}, fmt.Errorf("row.Scan - unable to retrieve new users id: %w", err)
	case err != nil:
		return entities.User{}, fmt.Errorf("row.Scan - database query failed %w: ", err)
	}
	return entities.CreateUser(id), nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id FROM users WHERE email=$1", email)
	var id int
	err := row.Scan(&id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return entities.User{}, fmt.Errorf("row.Scan - unable to retrieve new users id: %w", err)
	case err != nil:
		return entities.User{}, fmt.Errorf("row.Scan - database query failed %w: ", err)
	}
	return entities.CreateUser(id), nil
}

package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/LevinStrike/lux-backend/internal/core/apperror"
	"github.com/LevinStrike/lux-backend/internal/core/entities"
)

func (r *Repository) Login(ctx context.Context, username, password string) (entities.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, password FROM users WHERE email=$1;", username)
	var id int
	var dbpassword string
	err := row.Scan(&id, &dbpassword)
	switch {
	case err == sql.ErrNoRows:
		return entities.User{}, apperror.NewAuthenticationError(errors.New("invalid user credentials"))
	case err != nil:
		return entities.User{}, fmt.Errorf("database query failed because %s", err)
	}

	if err := ComparePassword(password, dbpassword); err != nil {
		return entities.User{}, apperror.NewAuthenticationError(err)
	}
	return entities.CreateUser(id), nil
}

func (r *Repository) SignUp(ctx context.Context, username, password string) (entities.User, error) {
	hash, err := HashPassword(password)
	if err != nil {
		return entities.User{}, fmt.Errorf("HashPassword: %W", err)
	}
	row := r.db.QueryRowContext(ctx, "INSERT INTO users(email, password) VALUES($1, $2) RETURNING id;", username, hash)
	var id int
	err = row.Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return entities.User{}, errors.New("unable to retrieve new users id")
	case err != nil:
		if strings.Contains(err.Error(), "duplicate key value violates") {
			return entities.User{}, apperror.NewBadRequestError(errors.New("user already exists"))
		}
		return entities.User{}, fmt.Errorf("database query failed because %s", err)
	}
	return entities.CreateUser(id), nil
}

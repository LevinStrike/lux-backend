package usecases

import (
	"context"
	"errors"

	"github.com/LevinStrike/lux-backend/internal/core/apperror"
	"github.com/LevinStrike/lux-backend/internal/core/entities"
)

var ErrInvalidUserCredentials = apperror.NewAuthenticationError(errors.New("unable to authenticate with invalid credentials"))

func (u *Usecases) Login(ctx context.Context, username, password string) (entities.User, error) {
	return u.repository.Login(ctx, username, password)
}

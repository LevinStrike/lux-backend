package usecases

import (
	"context"

	"github.com/LevinStrike/lux-backend/internal/core/entities"
)

func (u *Usecases) SignUp(ctx context.Context, username, password string) (entities.User, error) {
	return u.repository.SignUp(ctx, username, password)
}

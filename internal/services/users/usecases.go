package users

import (
	"context"

	"github.com/LevinStrike/lux-backend/internal/core/entities"
)

type Usecases interface {
	Login(ctx context.Context, username, password string) (entities.User, error)
	SignUp(ctx context.Context, username, password string) (entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
}

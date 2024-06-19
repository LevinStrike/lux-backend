package usecases

import (
	"context"

	"github.com/LevinStrike/lux-backend/internal/core/entities"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	// GetUserByID(ctx context.Context, id int) (entities.User, error)
	SignUp(ctx context.Context, username, password string) (entities.User, error)
	Login(ctx context.Context, username, password string) (entities.User, error)
}

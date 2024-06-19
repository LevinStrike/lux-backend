package usecases

import (
	"context"
	"fmt"

	"github.com/LevinStrike/lux-backend/internal/core/entities"
)

type Usecases struct {
	repository Repository
}

func NewUsecases(repository Repository) *Usecases {
	return &Usecases{
		repository: repository,
	}
}

func (u *Usecases) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	user, err := u.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return entities.User{}, fmt.Errorf("u.repository.GetUserByEmail: %w", err)
	}
	return user, nil
}

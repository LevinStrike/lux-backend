package users

import "github.com/LevinStrike/lux-backend/internal/services/users/usecases"

type Repository interface {
	usecases.Repository
}

package seed

import (
	"context"

	"github.com/LevinStrike/lux-backend/internal/services/users"
)

type Seeder struct {
	usersService *users.Service
}

type Services struct {
	Users *users.Service
}

func NewSeeder(services Services) *Seeder {
	return &Seeder{
		usersService: services.Users,
	}
}

func (s *Seeder) Seed(ctx context.Context) {
	s.usersService.SignUp(ctx, "bob", "securepassword1")
}

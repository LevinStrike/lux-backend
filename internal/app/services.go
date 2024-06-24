package app

import (
	"fmt"
	"log"
	"os"

	"github.com/LevinStrike/lux-backend/internal/repositories/postgres"
	"github.com/LevinStrike/lux-backend/internal/services/users"
)

type Services struct {
	Users *users.Service
}

func LoadServices(databases *databases) *Services {
	user, err := LoadUserService(databases)
	if err != nil {
		log.Fatal(fmt.Errorf("LoadAuthService: %w", err))
	}
	return &Services{Users: user}
}

func LoadUserService(databases *databases) (*users.Service, error) {
	var service *users.Service
	var err error
	switch os.Getenv("database_type") {
	default:
		service, err = users.NewService(nil,
			users.WithRepository(postgres.NewRepository(databases.Sql)),
		)
		if err != nil {
			return nil, fmt.Errorf("users.NewService: %w", err)
		}
		return service, err
	}
}

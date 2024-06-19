package users

import (
	"fmt"

	"github.com/LevinStrike/lux-backend/internal/services/users/usecases"
)

type Service struct {
	Usecases
	repository Repository
}

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func NewService(config *Config, configurations ...Configuration) (*Service, error) {
	s := &Service{}

	for _, configuration := range configurations {
		if err := configuration(s); err != nil {
			return nil, fmt.Errorf("configuration: %w", err)
		}
	}
	s.Usecases = usecases.NewUsecases(s.repository)
	return s, nil
}

func WithRepository(repository Repository) Configuration {
	return func(s *Service) error {
		s.repository = repository
		return nil
	}
}

type Configuration func(service *Service) error

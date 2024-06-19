package v1

import (
	"fmt"

	"github.com/LevinStrike/lux-backend/internal/core/apperror"
	"github.com/LevinStrike/lux-backend/internal/services/users"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	userService *users.Service
	router      chi.Router
}

type Configuration func(r *Router) error

func NewRouter(configurations ...Configuration) (*Router, error) {
	r := &Router{}
	for _, configuration := range configurations {
		if err := configuration(r); err != nil {
			return nil, fmt.Errorf("configuration: %w", err)
		}
	}
	return r, nil
}

func WithUserService(us *users.Service) Configuration {
	return func(r *Router) error {
		if us == nil {
			return apperror.ErrUnexpectedNillValues
		}
		r.userService = us
		return nil
	}
}

func WithRouter(router chi.Router) Configuration {
	return func(r *Router) error {
		if router == nil {
			return apperror.ErrUnexpectedNillValues
		}
		v1 := chi.NewRouter()
		router.Mount("/api/v1", v1)
		router = v1
		r.router = router
		return nil
	}
}

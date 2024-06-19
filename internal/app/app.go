//go:build !test

package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LevinStrike/lux-backend/internal/app/seed"
	v1 "github.com/LevinStrike/lux-backend/internal/controllers/rest/v1"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	_ "github.com/lib/pq"
)

func Start(ctx context.Context) {
	dbs := LoadDatabases(ctx)
	services := LoadServices(dbs)
	seeder := seed.NewSeeder(seed.Services{Users: services.Users})
	seeder.Seed(ctx)
	startHttp(services)
}

func startHttp(services *Services) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))
	v1Router, err := v1.NewRouter(
		v1.WithRouter(r),
		v1.WithUserService(services.Users),
	)
	if err != nil {
		log.Fatal(fmt.Errorf("v1.NewRouter: %w", err))
	}
	v1Router.AttachRoutes()

	if err := http.ListenAndServe(":8088", r); err != nil {
		panic(fmt.Errorf("http.ListenAndServe: %w", err))
	}
}

package app

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type databases struct {
	Sql *sql.DB
}

//go:embed migrations/*.sql
var postgresMigrations embed.FS

func LoadDatabases(ctx context.Context) *databases {
	switch os.Getenv("database_type") {
	default:
		return LoadPostgres(ctx)
	}
}

func LoadPostgres(ctx context.Context) *databases {
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATBASE_PORT")
	dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	wg := &sync.WaitGroup{}
	if os.Getenv("environment") != "prod" {
		wg.Add(1)
		go func() {
			postgresContainer, err := postgres.RunContainer(ctx,
				testcontainers.WithImage("docker.io/postgres:16-alpine"),
				postgres.WithDatabase(dbName),
				postgres.WithUsername(dbUser),
				postgres.WithPassword(dbPassword),
				testcontainers.WithWaitStrategy(
					wait.ForLog("database system is ready to accept connections").
						WithOccurrence(2).
						WithStartupTimeout(5*time.Second)),
			)
			if err != nil {
				log.Fatalf("failed to start container: %s", err)
			}
			dbConnectionString = postgresContainer.MustConnectionString(ctx) + "sslmode=disable"
			fmt.Println(dbConnectionString)
			wg.Done()

			defer func() {
				if err := postgresContainer.Terminate(ctx); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()
			<-ctx.Done()
		}()
	}

	wg.Wait()
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal(fmt.Errorf("sql.Open: %w", err))
	}

	if err = db.Ping(); err != nil {
		log.Fatal(fmt.Errorf("db.Ping: %w", err))
	}

	goose.SetBaseFS(postgresMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	return &databases{Sql: db}
}

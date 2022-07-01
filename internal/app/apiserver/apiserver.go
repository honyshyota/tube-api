package apiserver

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"

	pgstore "github.com/honyshyota/tube-api-go/internal/app/store/pg"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Start() error {
	db, err := newDB()
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	store := pgstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(os.Getenv("BIND_ADDR"), srv.router)
}

func newDB() (*sql.DB, error) {
	// DB connect
	dataBaseURL := os.Getenv("PG_URL")

	db, err := sql.Open("postgres", dataBaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if db != nil {
		logrus.Println("Running PostgreSQL migrations")
		if err := runPgMigrations(dataBaseURL); err != nil {
			return nil, fmt.Errorf("runPgMigrations failed: %w", err)
		}
		logrus.Println("Migrations done")
	}

	return db, nil
}

func runPgMigrations(pgURL string) error {
	MigrationsPath := os.Getenv("PG_MIGRATIONS_PATH")

	if MigrationsPath == "" {
		fmt.Println(1)
		return nil
	}

	if pgURL == "" {
		fmt.Println(2)
		return errors.New("no cfg.PgURL provided")
	}

	m, err := migrate.New(
		MigrationsPath,
		pgURL,
	)
	if err != nil {
		fmt.Println(3)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(4)
		return err
	}

	return nil
}

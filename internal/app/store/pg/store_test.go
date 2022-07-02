package pgstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("PG_TEST_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:postgres@localhost:5432/postgres_test?sslmode=disable"
	}

	os.Exit(m.Run())
}

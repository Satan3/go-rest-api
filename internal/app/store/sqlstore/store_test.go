package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=go-rest-api-test user=postgres password=postgres sslmode=disable"
	}
}

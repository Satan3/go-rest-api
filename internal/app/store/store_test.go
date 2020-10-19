package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(t *testing.M) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=go-rest-api-test user=postgres password=postgres sslmode=disable"
	}
}

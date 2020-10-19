.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./ ...

.PHONY: migrate
migrate:
	migrate -path migrations -database postgres://localhost/go-rest-api?sslmode=disable&user=postgres&password=postgres up

.DEFAULT_GOAL := build
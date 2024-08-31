include .env

dev:
	@air -c .air.toml

build:
	@go build -o ./bin/main ./cmd/api/main.go

run: build
	@./bin/main

install.goose:
	@go install github.com/pressly/goose/v3/cmd/goose@latest

goose.up:
	@cd sql/migrations && goose postgres ${DB_URL} up

goose.down:
	@cd sql/migrations && goose postgres ${DB_URL} down
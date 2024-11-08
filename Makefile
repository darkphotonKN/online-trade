# Load the .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Construct the DB_STRING dynamically
DB_STRING=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

build:
	@go build -o bin/online-shop ./cmd/

run: build
	@./bin/online-shop

# Run tests with verbose output and coverage
test:
	@go test -v ./... -cover

# Run tests with coverage output and preview in a browser
test-preview:
	@go test ./filename/ -coverprofile=coverage.out 
	@go tool cover -html=coverage.out

# Migration commands using Goose
migrate-up:
	@goose -dir ./migrations postgres "$(DB_STRING)" up

migrate-down:
	@goose -dir ./migrations postgres "$(DB_STRING)" down

migrate-status:
	@goose -dir ./migrations postgres "$(DB_STRING)" status

migrate-down-to:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make migrate-down-to VERSION=<version>"; \
		exit 1; \
	fi; \
	goose -dir ./migrations postgres "$(DB_STRING)" down-to $(VERSION)

migrate-reset:
	@goose -dir ./migrations postgres "$(DB_STRING)" reset

.PHONY: run test migrate-up migrate-down migrate-status









BINARY_NAME=language-learning-app

.PHONY: all build run test clean help migrate

all: build

build:
	@swag init
	@echo "Building the application..."
	@mkdir -p build
	@go build -o build/$(BINARY_NAME) cmd/api/main.go
	@go build -o build/migrate cmd/migrate/main.go

run: build
	@echo "Running the application..."
	@./build/$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@rm -f build/$(BINARY_NAME)
	@rm -f build/migrate

# Migrations
migrate:
	@echo "Running database migrations..."
	@go run cmd/migrate/main.go

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build    Build the application"
	@echo "  run      Run the application"
	@echo "  test     Run the tests"
	@echo "  clean    Clean the build artifacts"
	@echo "  migrate  Run database migrations"
	@echo "  help     Display this help message"

.DEFAULT_GOAL := help

.PHONY: run init build watch migrate generate
	
include .env

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

run:
	@go run cmd/api/main.go

init:
	@if command -v goimports > /dev/null; then \
		go mod tidy; \
		goimports -w .; \
		rm -rf .git; \
		git init; \
	    echo "Initializing...";\
	else \
		go install golang.org/x/tools/cmd/goimports@latest; \
		go mod tidy; \
		goimports -w .; \
		rm -rf .git; \
		git init; \
	    echo "Initializing...";\
	fi

watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

generate:
	@if command -v sqlc > /dev/null; then \
		sqlc generate; \
	else \
	    read -p "SQLc is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
			sqlc generate; \
	    else \
	        echo "You chose not to install SQLc. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

migrate:
	@psql "$(POSTGRESQL_URL)" -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"
	POSTGRESQL_URL="$(POSTGRESQL_URL)" tern migrate --migrations internal/database/sql/migrations/

.DEFAULT_GOAL := help

DATABASE_URL ?= postgres://madagascar:madagascar@localhost:5432/madagascar?sslmode=disable
export DATABASE_URL

.PHONY: help
help: ## List the available targets
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN{FS=":.*?## "};{printf "  %-12s %s\n", $$1, $$2}'

.PHONY: run
run: ## Run the server
	go run ./cmd/server

.PHONY: build
build: ## Build a single static binary into bin/
	CGO_ENABLED=0 go build -trimpath -ldflags='-s -w' -o bin/server ./cmd/server

.PHONY: test
test: ## Run the tests
	go test ./...

.PHONY: cover
cover: ## Run the tests with a coverage summary
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out | tail -1

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: fmt
fmt: ## Format the code
	go fmt ./...

.PHONY: check
check: fmt vet test ## Format, vet and test

.PHONY: db-up
db-up: ## Start PostgreSQL
	docker compose up -d

.PHONY: db-down
db-down: ## Stop PostgreSQL
	docker compose down

.PHONY: db-shell
db-shell: ## Open psql against the local database
	docker compose exec postgres psql -U madagascar -d madagascar

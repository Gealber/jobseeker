.DEFAULT_GOAL := help
.PHONY : build

DATABASE_DSN=$(DATABASE_DSN)
ifeq ($(ENV),test)
	DATABASE_DSN="postgresql://root@gulolio:26257/defaultdb?sslmode=disable"
endif

run: ## Run code
	@go run main.go

build: ## Build binary
	@mkdir -p bin
	@go build  -o bin/node main.go

test: ## Run tests
	@go test -v -race ./... -coverprofile=coverage.out

migration-up: ## Migartion up
	@goose -dir=./database/migrations postgres $(DATABASE_DSN) up

migration-down: ## Migartion down
	@goose -dir=./database/migrations postgres $(DATABASE_DSN) dow

help: ## Show commands availables
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
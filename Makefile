.PHONY: help

.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

vet: ## Run vet
	go vet ./...

fmt: ## Run gofmt
	gofmt -s -w .

build: ## Build local environment
	docker context use default
	docker compose -f docker-compose-dev.yml build

up: ## Setup local environment
	docker context use default
	docker compose -f docker-compose-dev.yml up

down: ## Teardown local environment
	docker context use default
	docker compose -f docker-compose-dev.yml down

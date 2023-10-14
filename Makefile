# ================================================
# VARIABLES
# ================================================
ENV_DEV_FILE = --env-file .env.development

MKDOCS_SERVICE = mkdocs
MKDOCS_CONTAINER = kafrest-mkdocs

# ================================================
# COMMON
# ================================================
# default: local-stack run-api-container run-worker-container ## Build and run all containers

help: ## Print available commands
	$(info ========================================)
	$(info Available Commands:)
	@grep '^[[:alnum:]_-]*:.* ##' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS=":.* ## "}; {printf "make %-25s %s\n", $$1, $$2};'
	$(info ========================================)
.PHONY: help

lint: ## Run linter
	docker run -t --rm -v .:/app -w /app golangci/golangci-lint:v1.54.2 golangci-lint run -v
.PHONY: lint

dev: ## Run the app in development mode
	APP_ENV=development go run cmd/api/main.go
.PHONY: dev

up:
	@docker-compose up --build -d 
.PHONY: up

stop:
	@docker-compose stop
.PHONY: stop

run-docs: ## Build and run mkdocs
	@docker-compose $(ENV_DEV_FILE) up -d --build $(MKDOCS_SERVICE)
.PHONY: build-docs

# =========================================
# MKDOCS
# =========================================
clear-docs:
	@docker-compose stop $(MKDOCS_SERVICE)
	@docker rm $(MKDOCS_CONTAINER) -v
.PHONY: clear-docs

logs-docs:
	@docker logs -f $(MKDOCS_CONTAINER)
.PHONY: logs-docs
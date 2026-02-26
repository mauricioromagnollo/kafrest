# ================================================
# VARIABLES
# ================================================
APP_NAME = kafrest
ENV_FILE = --env-file .env

APP_SERVICE_NAME = app
APP_CONTAINER_NAME= $(APP_NAME)_app

ZOOKEEPER_SERVICE_NAME = zookeeper
ZOOKEEPER_CONTAINER_NAME = $(APP_NAME)_zookeeper

KAFKA_SERVICE_NAME = kafka
KAFKA_CONTAINER_NAME = $(APP_NAME)_kafka

AKHQ_SERVICE_NAME = akhq
AKHQ_CONTAINER_NAME = $(APP_NAME)_akhq

SCHEMA_REGISTRY_SERVICE_NAME = schema-registry
SCHEMA_REGISTRY_CONTAINER_NAME = $(APP_NAME)_schema-registry

KAFKA_CREATE_TOPICS_SERVICE_NAME = kafka-create-topics
KAFKA_CREATE_TOPICS_CONTAINER_NAME = $(APP_NAME)_kafka-create-topics

MKDOCS_SERVICE_NAME = mkdocs
MKDOCS_CONTAINER_NAME = $(APP_NAME)_mkdocs

AKHQ_SERVICE = akhq
AKHQ_CONTAINER = $(APP_NAME)_akhq

# ================================================
# COMMANDS
# ================================================
default: zookeeper kafka schema-registry topics app

all: zookeeper kafka schema-registry topics app akhq docs ## Build all "make" services including: AKHQ, MKDOCS
.PHONY: all

help: ## Print all available commands
	$(info ========================================)
	$(info Available Commands:)
	@grep '^[[:alnum:]_-]*:.* ##' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS=":.* ## "}; {printf "make %-25s %s\n", $$1, $$2};'
	$(info ========================================)
.PHONY: help

up: ## Up all containers in detached mode
	@docker compose up -d
.PHONY: up

down: ## Down all containers, but keep images, networks, and volumes
	@docker compose down
.PHONY: down

stop: ## Stop all containers, but keep images, networks, and volumes
	@docker compose stop
.PHONY: stop

clear: ## Remove all containers, images, networks, and volumes
	@docker compose down --rmi all --volumes --remove-orphans
.PHONY: clear

# =========================================
# KAFKA
# =========================================
kafka: ## Build and run the kafka container
	@docker compose $(ENV_FILE) up -d --build --remove-orphans $(KAFKA_SERVICE_NAME)
.PHONY: kafka

topics: ## Build container only for creating Kafka topics, drop it after the topics are created
	@docker compose $(ENV_FILE) up -d --build --remove-orphans $(KAFKA_CREATE_TOPICS_SERVICE_NAME)
.PHONY: topics

zookeeper: ## Build and run the zookeeper container
	@docker compose $(ENV_FILE) up -d --build --remove-orphans $(ZOOKEEPER_SERVICE_NAME)
.PHONY: zookeeper

# =========================================
# APP
# =========================================
app: ## Build and run the app worker
	@docker compose $(ENV_FILE) up -d --build $(APP_SERVICE_NAME)
.PHONY: app

open: ## Open a shell session inside the app container
	@docker exec -it $(APP_CONTAINER_NAME) /bin/sh
.PHONY: open

logs: ## Show app logs
	@docker logs -f $(APP_CONTAINER_NAME)
.PHONY: logs

lint: ## Run linter
	docker run -t --rm -v ${PWD}/:/app -w /app golangci/golangci-lint:v2.9.0 golangci-lint run -v
.PHONY: lint

test: ## Run all app tests
	@echo "Running tests..."
	go test -v ./...
.PHONY: test

test-coverage: ## Run all app test coverage
	go test -v -race -coverprofile=cover.out ./...
	go tool cover -func=cover.out
.PHONY: test-coverage

test-coverage-web: ## Run test coverage and show in browser
	go test -v -race -coverprofile=cover.out ./... && go tool cover -html=cover.out
.PHONY: test-coverage-web

test-race: # Run data race tests
	CGO_ENABLED=1 go test -race ./...
.PHONY: test-race

# =========================================
# MKDOCS
# =========================================
docs: ## Build and run mkdocs
	@docker-compose $(ENV_FILE) up -d --build $(MKDOCS_SERVICE_NAME)
.PHONY: docs

docs-clear: ## Remove mkdocs Container and Image
	@docker-compose stop $(MKDOCS_SERVICE_NAME)
	@docker rm $(MKDOCS_CONTAINER_NAME) -v
.PHONY: docs-clear

docs-logs:
	@docker logs -f $(MKDOCS_CONTAINER_NAME)
.PHONY: docs-logs

# =====================================
# AKHQ
# =====================================
akhq: ## Build akhq container
	@docker compose $(ENV_FILE) up -d --build --remove-orphans $(AKHQ_SERVICE_NAME)
.PHONY: akhq

akhq-logs: ## Show AKHQ Logs
	@docker logs -f $(AKHQ_CONTAINER)
.PHONY: akhq-logs

akhq-clear: ## Remove AKHQ Container and Image
	@docker stop $(AKHQ_CONTAINER)
	@docker rm $(AKHQ_CONTAINER)
.PHONY: akhq-clear

# ================================================
# VARIABLES
# ================================================
ENV_FILE = --env-file .env

API_SERVICE = api
API_CONTAINER = kafrest_api

KAFKA_SERVICE = kafka
KAFKA_CONTAINER = kafrest_kafka

KAFKA_CREATE_TOPICS_SERVICE = kafka-create-topics
KAFKA_CREATE_TOPICS_CONTAINER = kafrest_kafka-create-topics

MKDOCS_SERVICE = mkdocs
MKDOCS_CONTAINER = kafrest_mkdocs

KAFREST_LATEST_SERVICE = kafrest-latest
KAFREST_LATEST_CONTAINER = kafrest_latest

AKHQ_SERVICE = akhq
AKHQ_CONTAINER = kafrest_akhq

# ================================================
# COMMON
# ================================================
default: zookeeper kafka kafka-create-topics akhq api ## Build and run all containers

help: ## Print available commands
	$(info ========================================)
	$(info Available Commands:)
	@grep '^[[:alnum:]_-]*:.* ##' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS=":.* ## "}; {printf "\t%-25s %s\n", $$1, $$2};'
	$(info ========================================)
	$(info make <command>)
	$(info )
.PHONY: help

lint: ## Run linter
	docker run -t --rm -v .:/app -w /app golangci/golangci-lint:v1.54.2 golangci-lint run -v
.PHONY: lint

stop: ## Stop all containers
	@docker-compose $(ENV_FILE) stop
.PHONY: stop

start: ## Start all containers
	@docker-compose $(ENV_FILE) start
.PHONY: start

clear: ## Stop containers, remove images, networks, and volumes
	@docker compose down --rmi all --volumes --remove-orphans
.PHONY: clear

# =========================================
# API
# =========================================
api: ## Build and run API
	@docker-compose $(ENV_FILE) up -d --build $(API_SERVICE)
.PHONY: api

logs: ## Show API Logs
	@docker logs -f $(API_CONTAINER)
.PHONY: logs

api-clear: ## Remove API Container and Image
	@docker-compose stop $(API_SERVICE)
	@docker rm $(API_CONTAINER) -v
.PHONY: api-clear

# =========================================
# MKDOCS
# =========================================
docs: ## Build and run mkdocs
	@docker-compose $(ENV_FILE) up -d --build $(MKDOCS_SERVICE)
.PHONY: docs

docs-clear: ## Remove mkdocs Container and Image
	@docker-compose stop $(MKDOCS_SERVICE)
	@docker rm $(MKDOCS_CONTAINER) -v
.PHONY: docs-clear

docs-logs:
	@docker logs -f $(MKDOCS_CONTAINER)
.PHONY: docs-logs

# =========================================
# KAFKA
# =========================================
kafka:
	@docker-compose $(ENV_FILE) up -d --build $(KAFKA_SERVICE)
.PHONY: kafka

kafka-create-topics:
	@docker-compose $(ENV_FILE) up -d --build $(KAFKA_CREATE_TOPICS_SERVICE)
.PHONY: kafka-create-topics

zookeeper:
	@docker-compose $(ENV_FILE) up -d --build $(ZOOKEEPER_SERVICE)
.PHONY: zookeeper

# =====================================
# AKHQ
# =====================================
akhq: ## Build akhq container
	@docker compose $(ENV_FILE) up -d --build $(AKHQ_SERVICE)
.PHONY: akhq

akhq-logs: ## Show AKHQ Logs
	@docker logs -f $(AKHQ_CONTAINER)
.PHONY: akhq-logs

akhq-clear: ## Remove AKHQ Container and Image
	@docker stop $(AKHQ_CONTAINER)
	@docker rm $(AKHQ_CONTAINER)
.PHONY: akhq-clear
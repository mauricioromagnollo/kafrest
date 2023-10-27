# ================================================
# VARIABLES
# ================================================
ENV_DEV_FILE = --env-file .env.development

MKDOCS_SERVICE = mkdocs
MKDOCS_CONTAINER = kafrest-mkdocs

API_SERVICE = kafrest
API_CONTAINER = kafrest

KAFKA_CREATE_TOPICS_SERVICE = kafka-create-topics
KAFKA_CREATE_TOPICS_CONTAINER = kafrest-kafka-create-topics

# ================================================
# COMMON
# ================================================
default: build-zookeeper build-kafka kafka-create-topics run-api ## Build and run all containers

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

stop: ## Stop all containers
	@docker-compose $(ENV_DEV_FILE) stop
.PHONY: stop

start: ## Start all containers
	@docker-compose $(ENV_DEV_FILE) start
.PHONY: start

clear: ## Remove all containers and volumes associated with this project
	@docker-compose $(ENV_DEV_FILE) down -v
.PHONY: clear

kafka-create-topics:
	@docker-compose $(ENV_DEV_FILE) up -d --build $(KAFKA_CREATE_TOPICS_SERVICE)
.PHONY: kafka-create-topics

# =========================================
# API
# =========================================
run-api: ## Build and run api
	@docker-compose $(ENV_DEV_FILE) up -d --build $(API_SERVICE)
.PHONY: run-api

logs-api: ## Show Oculus API Logs
	@docker logs -f $(API_CONTAINER)
.PHONY: api-logs

clear-api: ## Remove Oculus API Container and Image
	@docker-compose stop $(API_SERVICE)
	@docker rm $(API_CONTAINER) -v
.PHONY: api-clear

# =========================================
# MKDOCS
# =========================================
run-docs: ## Build and run mkdocs
	@docker-compose $(ENV_DEV_FILE) up -d --build $(MKDOCS_SERVICE)
.PHONY: build-docs

clear-docs:
	@docker-compose stop $(MKDOCS_SERVICE)
	@docker rm $(MKDOCS_CONTAINER) -v
.PHONY: clear-docs

logs-docs:
	@docker logs -f $(MKDOCS_CONTAINER)
.PHONY: logs-docs

# =========================================
# ZOOKEEPER
# =========================================
build-zookeeper:
	@docker-compose $(ENV_DEV_FILE) up -d --build $(ZOOKEEPER_SERVICE)
.PHONY: build-zookeeper

# =========================================
# KAFKA
# =========================================
build-kafka:
	@docker-compose $(ENV_DEV_FILE) up -d --build $(KAFKA_SERVICE)
.PHONY: build-kafka

# =========================================
# KAFDROP
# =========================================
build-kafdrop:
	@docker-compose $(ENV_DEV_FILE) up -d --build $(KAFDROP_SERVICE)
.PHONY: build-kafdrop

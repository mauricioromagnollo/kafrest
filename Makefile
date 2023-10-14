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
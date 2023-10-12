dev: ## Run the app in development mode
	APP_ENV=development go run cmd/api/main.go
.PHONY: dev

up:
	@docker-compose up --build -d 
.PHONY: up

stop:
	@docker-compose stop
.PHONY: stop
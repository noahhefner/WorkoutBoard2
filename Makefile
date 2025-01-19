DEV_COMPOSE_FILE=docker-compose-dev.yml
DEV_CMD=docker-compose -f $(DEV_COMPOSE_FILE)

# Default target
.PHONY: help
help:
	@echo "Makefile for Docker Compose"
	@echo "Usage:"
	@echo "  make up          - Start services in the background"
	@echo "  make down        - Stop services and remove containers"
	@echo "  make logs        - View logs of services"
	@echo "  make ps          - List running services"
	@echo "  make setup       - Install npm packages"
	@echo "  make format      - Run Prettier"

.PHONY: up
up:
	$(DEV_CMD) up --build --no-recreate -d

.PHONY: down
down:
	$(DEV_CMD) down

.PHONY: logs
logs:
	$(DEV_CMD) logs -f

.PHONY: ps
ps:
	$(DEV_CMD) ps

.PHONY: setup
setup:
	npm install

.PHONY: format
format:
	npx prettier . --write
.PHONY: help build up down dev logs clean test

# Default target
help:
	@echo "Available commands:"
	@echo "  build     - Build all Docker images"
	@echo "  up        - Start all services (production)"
	@echo "  dev       - Start development environment"
	@echo "  down      - Stop all services"
	@echo "  logs      - Show logs from all services"
	@echo "  clean     - Remove all containers and volumes"
	@echo "  test      - Run backend tests"

# Build all images
build:
	docker-compose build

# Start production environment
up:
	docker-compose up -d

# Start development environment
dev:
	docker-compose -f docker-compose.dev.yml up

# Stop all services
down:
	docker-compose down
	docker-compose -f docker-compose.dev.yml down

# Show logs
logs:
	docker-compose logs -f

# Clean everything
clean:
	docker-compose down -v --remove-orphans
	docker-compose -f docker-compose.dev.yml down -v --remove-orphans
	docker system prune -f

# Run backend tests
test:
	docker-compose exec backend go test ./...
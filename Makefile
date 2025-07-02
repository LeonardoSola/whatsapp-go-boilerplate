.PHONY: help build run stop clean logs shell

# Default target
help:
	@echo "Available commands:"
	@echo "  build    - Build the Docker image"
	@echo "  run      - Run the WhatsApp bot container"
	@echo "  stop     - Stop the container"
	@echo "  clean    - Remove container and image"
	@echo "  logs     - Show container logs"
	@echo "  shell    - Access container shell"
	@echo "  dev      - Run in development mode with volume mounting"

# Build the Docker image
build:
	docker-compose build

# Run the container
run:
	docker-compose up -d

# Run in development mode (with source code mounted)
dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up

# Stop the container
stop:
	docker-compose down

# Remove container and image
clean:
	docker-compose down --rmi all --volumes --remove-orphans

# Show container logs
logs:
	docker-compose logs -f

# Access container shell
shell:
	docker-compose exec whatsapp-bot sh

# Build and run in one command
build-and-run: build run
	@echo "Container is running. Use 'make logs' to view logs." 
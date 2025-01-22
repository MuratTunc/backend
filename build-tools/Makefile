# Environment file and variables
ENV_FILE := .env

##Stop any running Docker containers.
##Clean up any existing binaries.
##Build the services.
##Check if the necessary ports are free.
##Rebuild and start the Docker images.
##Fetch logs for all services.
##Clean up the built binaries.

## up_build: stops docker-compose (if running), builds all projects, checks ports, and starts docker-compose
up_build: down clean build_main_service
	@echo "Checking if service ports are available..."
	$(MAKE) check_ports
	@echo "Building (when required) and starting docker images with environment variables..."
	docker-compose up --build -d
	@echo "Docker images built and started!"
	@echo "Fetching logs for all services..."
	docker-compose logs --tail=50
	@echo "Cleaning up built binaries..."
	$(MAKE) clean
	@echo "up_build process completed!"

## check_ports: Checks if ports are available
check_ports:
	@echo "Checking if service ports are available..."
	@set -a; . ./$(ENV_FILE); set +a; \
		echo "Main Service Binary: $$MAIN_SERVICE_BINARY";
		if lsof -i :$$MAIN_SERVICE_PORT; then \
			echo "Error: One or more ports are already in use."; \
			exit 1; \
		else \
			echo "Ports are available."; \
		fi

## down: stop docker-compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## clean: removes built binaries
clean:
	@echo "Cleaning up built binaries..."
	@set -a; . ./$(ENV_FILE); set +a; \
		echo "MAIN_SERVICE_BINARY: $$MAIN_SERVICE_BINARY";
		[ -f ../main-service/$$MAIN_SERVICE_BINARY ] && rm -f ../main-service/$$MIN_SERVICE_BINARY || echo "File $$MAIN_SERVICE_BINARY does not exist";
	@echo "Done!"


## build_main_service: builds the main-service binary as a linux executable
build_main_service:
	@echo "Building main-service binary..."
	@set -a; . ./$(ENV_FILE); set +a; \
		cd ../main-service && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $$GOFULLPATH build -o $$MAIN_SERVICE_BINARY ./cmd/api
	@echo "Done!"


## logs: shows logs from all services
logs:
	@echo "Fetching logs for all services..."
	docker-compose logs -f

## help: displays the list of available commands
help:
	@grep -E '^##' $(MAKEFILE_LIST) | sed -e 's/## //'

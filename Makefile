# Define variables
DOCKER_COMPOSE = docker compose
COMPOSE_FILE = docker-compose.yml
BUILD_ARGS = --build

# Define targets

# Target to build and start the application
up:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up $(BUILD_ARGS)

# Target to start the application without rebuilding
start:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up

# Target to stop and remove containers
down:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down

# Target to stop the containers without removing them
stop:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) stop

# Target to remove stopped containers, networks, and volumes
clean:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down --volumes --remove-orphans

# Target to view logs of the application
logs:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) logs -f

# Target to rebuild the images without starting containers
rebuild:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) build

build:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) build --no-cache

.PHONY: up start down stop clean logs rebuild

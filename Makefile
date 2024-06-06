# Устанавливаем переменные
APP_NAME = bookstore
DOCKER_COMPOSE_FILE = docker-compose.yml

# Цель по умолчанию
.PHONY: all
all: build run

# Сборка приложения
.PHONY: build
build:
	@echo "Building the application..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

# Запуск приложения
.PHONY: run
run:
	@echo "Running the application..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Остановка приложения
.PHONY: stop
stop:
	@echo "Stopping the application..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Очистка сборки
.PHONY: clean
clean:
	@echo "Cleaning up..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down -v

.PHONY: test
test:
	@echo "Make is working!"
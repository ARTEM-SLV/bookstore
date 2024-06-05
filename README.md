# Bookstore API

Это REST API для управления коллекцией книг и авторами, которые их пишут, созданное на языке Go и использующее PostgreSQL для хранения данных. Приложение контейнеризировано с помощью Docker.

## Структура проекта

- `cmd/main.go`: Главный файл приложения.
- `pkg/handlers`: Обработчики HTTP-запросов.
- `pkg/services`: Логика работы с базой данных.
- `pkg/repositories`: Описание базы данных
- `migrations`: SQL файлы для миграции базы данных.
- `Dockerfile`: Инструкции для сборки Docker-образа.
- `docker-compose.yml`: Конфигурация Docker Compose.

## Требования

- Go 1.22
- Docker
- Docker Compose

## Инструкции по сборке и запуску

### Сборка и запуск

Используйте `Makefile` для автоматизации процесса сборки и запуска.

1. Сборка приложения:
    ```sh
    make build
    ```

2. Запуск приложения:
    ```sh
    make run
    ```

3. Остановка и удаление контейнеров:
    ```sh
    make down
    ```

### Проверка состояния контейнеров

Для проверки состояния контейнеров используйте команду:
```sh
docker-compose ps

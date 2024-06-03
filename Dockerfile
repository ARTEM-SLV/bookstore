# Используем конкретную версию официального изображения Go в качестве базового
FROM golang:1.22

## Устанавливаем утилиту migrate
#RUN apk add --no-cache curl \
#  && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz \
#  && mv migrate.linux-amd64 /usr/local/bin/migrate \
#  && apk del curl

# Устанавливаем зависимости
RUN apk add --no-cache git

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum файлы в контейнер
COPY go.mod go.sum ./

# Загружаем все зависимости. Кэшируется, если go.mod и go.sum не изменяются
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

# Сборка Go-приложения
RUN go build -o main ./cmd

## Копируем файлы миграций
#COPY ./repositories/migrations ./migrations

## Выполняем миграции
#CMD ["sh", "-c", "migrate -path ./repositories/migrations -database $DATABASE_URL up && ./main"]

# Устанавливаем команду для запуска приложения
CMD ["./main"]

# Открываем порт 8080 для доступа извне
EXPOSE 8080

## Команда для запуска миграций и приложения
#CMD migrate -path /app/migrations -database "postgres://postgres:postgres@db:5432/postgres?sslmode=disable" up && ./main
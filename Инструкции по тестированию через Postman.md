# Инструкции по тестированию через Postman

## Создание коллекции:

Создайте новую коллекцию в Postman и назовите её Bookstore API.
Создание запросов:

В коллекции создайте запросы для каждого из описанных выше методов.
Примеры запросов:

```sh
POST /books:
Метод: POST
URL: http://localhost:8080/books
Тело запроса: JSON (как указано выше)
```

```sh
GET /books:
Метод: GET
URL: http://localhost:8080/books
```

```sh
GET /books/{id}:
Метод: GET
URL: http://localhost:8080/books/1
```

```sh
PUT /books/{id}:
Метод: PUT
URL: http://localhost:8080/books/1
Тело запроса: JSON (как указано выше)
```

```sh
DELETE /books/{id}:
Метод: DELETE
URL: http://localhost:8080/books/1
```

```sh
POST /authors:
Метод: POST
URL: http://localhost:8080/authors
Тело запроса: JSON (как указано выше)
```

```sh
GET /authors:
Метод: GET
URL: http://localhost:8080/authors
```

```sh
GET /authors/{id}:
Метод: GET
URL: http://localhost:8080/authors/1
```

```sh
PUT /authors/{id}:
Метод: PUT
URL: http://localhost:8080/authors/1
Тело запроса: JSON (как указано выше)
```

```sh
DELETE /authors/{id}:
Метод: DELETE
URL: http://localhost:8080/authors/1
```

PUT /books/{book_id}/authors/{author_id}:
Метод: PUT
URL: http://localhost:8080/books/1/authors/1
Тело запроса: JSON (как указано выше)

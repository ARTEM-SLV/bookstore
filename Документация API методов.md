# HTTP Методы

## 1. Книги

### POST /books
Описание: Добавляет новую книгу в коллекцию.

Тело запроса:
```sh
{
    "title": "Название книги",
    "author_id": 1,
    "year": 2023,
    "isbn": "123-4567890123"
}
```

Ответ:
201 Created: Книга успешно добавлена.
400 Bad Request: Некорректные данные в запросе.

### GET /books
Описание: Получает список всех книг в коллекции.

Ответ:
200 OK: Возвращает массив книг.
500 Internal Server Error: Ошибка на сервере.

### GET /books/{id}
Описание: Получает информацию о книге по её уникальному идентификатору.

Ответ:
200 OK: Возвращает данные книги.
404 Not Found: Книга с указанным идентификатором не найдена.

### PUT /books/{id}
Описание: Обновляет информацию о книге по её уникальному идентификатору.

Тело запроса:
```sh
{
    "title": "Новое название книги",
    "author_id": 1,
    "year": 2024,
    "isbn": "123-4567890123"
}
```

Ответ:
200 OK: Книга успешно обновлена.
400 Bad Request: Некорректные данные в запросе.
404 Not Found: Книга с указанным идентификатором не найдена.

### DELETE /books/{id}
Описание: Удаляет книгу по её уникальному идентификатору.

Ответ:
204 No Content: Книга успешно удалена.
404 Not Found: Книга с указанным идентификатором не найдена.

## 2. Авторы

### POST /authors
Описание: Добавляет нового автора в коллекцию.

Тело запроса:
```sh
{
    "first_name": "Имя",
    "last_name": "Фамилия",
    "bio": "Краткая биография",
    "birth_date": "1976-08-11"
}
```

Ответ:
201 Created: Автор успешно добавлен.
400 Bad Request: Некорректные данные в запросе.

### GET /authors
Описание: Получает список всех авторов в коллекции.

Ответ:
200 OK: Возвращает массив авторов.
500 Internal Server Error: Ошибка на сервере.

### GET /authors/{id}
Описание: Получает информацию об авторе по его уникальному идентификатору.

Ответ:
200 OK: Возвращает данные автора.
404 Not Found: Автор с указанным идентификатором не найден.

### PUT /authors/{id}
Описание: Обновляет информацию об авторе по его уникальному идентификатору.

Тело запроса:
```sh
{
    "first_name": "Новое имя",
    "last_name": "Новая фамилия",
    "bio": "Обновленная биография",
    "birth_date": "1980-01-01"
}
```

Ответ:
200 OK: Автор успешно обновлен.
400 Bad Request: Некорректные данные в запросе.
404 Not Found: Автор с указанным идентификатором не найден.

### DELETE /authors/{id}
Описание: Удаляет автора по его уникальному идентификатору.

Ответ:
204 No Content: Автор успешно удален.
404 Not Found: Автор с указанным идентификатором не найден.

## 3. Транзакционное обновление

### PUT /books/{book_id}/authors/{author_id}
Описание: Одновременно обновляет сведения о книге и авторе.

Тело запроса:
```sh
{
    "book": {
        "title": "Новое название книги",
        "author_id": 2,
        "year": 2024,
        "isbn": "123-4567890123"
    },
    "Author": {
        "first_name": "Новое имя",
        "last_name": "Новая фамилия",
        "bio": "Обновленная биография",
        "birth_date": "1980-01-01"
    }
}
```

Ответ:
200 OK: Книга и автор успешно обновлены.
400 Bad Request: Некорректные данные в запросе.
404 Not Found: Книга или автор с указанными идентификаторами не найдены.

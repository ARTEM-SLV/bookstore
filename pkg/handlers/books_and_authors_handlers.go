package handlers

import (
	models2 "bookstore/models"
	"bookstore/pkg/repositories/postgre"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strconv"
)

// Структура для обновления книги и автора
type UpdateRequest struct {
	Book   models2.Book              `json:"book"`
	Author models2.AuthorWithoutTime `json:"author"`
}

func UpdateBookAndAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["book_id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	authorID, err := strconv.Atoi(vars["author_id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	var updateRequest UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updateRequest.Book.ID != bookID || updateRequest.Author.ID != authorID {
		http.Error(w, "Mismatched book or author ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	tx, err := postgre.GetDB().Begin(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "UPDATE books SET title=$1, author_id=$2, year=$3, isbn=$4 WHERE id=$5",
		updateRequest.Book.Title, updateRequest.Book.AuthorID, updateRequest.Book.Year, updateRequest.Book.ISBN, bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec(ctx, "UPDATE authors SET first_name=$1, last_name=$2, biography=$3, birth_date=$4 WHERE id=$5",
		updateRequest.Author.FirstName, updateRequest.Author.LastName, updateRequest.Author.Biography, updateRequest.Author.BirthDate, authorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := "Book and Author updated"
	w.Write([]byte(response))
}

// GetAuthorAndBooks handles GET /author_with_books/{id}
func GetAuthorAndBooks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	conn, err := postgre.GetDB().Acquire(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	var author models2.Author
	query := "SELECT id, first_name, last_name, biography, birth_date FROM authors WHERE id=$1"
	err = conn.QueryRow(r.Context(), query, id).Scan(&author.ID, &author.FirstName, &author.LastName, &author.Biography, &author.BirthDate)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Author not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	query = `SELECT id, title, year, isbn FROM books WHERE author_id=$1`
	rows, err := conn.Query(r.Context(), query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models2.BookWithAuthor
	for rows.Next() {
		var book models2.BookWithAuthor
		err := rows.Scan(&book.ID, &book.Title, &book.Year, &book.ISBN)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
	json.NewEncoder(w).Encode(books)
}

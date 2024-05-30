package handlers

import (
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strconv"

	"bookstore/db"
	"bookstore/pkg/models"
	"github.com/gorilla/mux"
)

// CreateBook handles POST /books
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := db.GetDB().Acquire(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	sql := `INSERT INTO books (title, author_id, year, isbn) VALUES ($1, $2, $3, $4) RETURNING id`
	err = conn.QueryRow(r.Context(), sql, book.Title, book.AuthorID, book.Year, book.ISBN).Scan(&book.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GetBooks handles GET /books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	conn, err := db.GetDB().Acquire(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	query := `SELECT b.id, b.title, CONCAT(a.first_name, ' ', a.last_name) as author, b.year, b.isbn FROM public.books as b INNER JOIN public.authors as a on b.author_id = a.id`
	rows, err := conn.Query(r.Context(), query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.BookWithAuthor
	for rows.Next() {
		var book models.BookWithAuthor
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.ISBN)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook handles GET /books/{id}
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	conn, err := db.GetDB().Acquire(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	var book models.Book
	sql := "SELECT id, title, author_id, year, isbn FROM books WHERE id=$1"
	err = conn.QueryRow(r.Context(), sql, id).Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year, &book.ISBN)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBook handles PUT /books/{id}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book.ID = id

	conn, err := db.GetDB().Acquire(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	sql := "UPDATE books SET title=$1, author_id=$2, year=$3, isbn=$4 WHERE id=$5"
	_, err = conn.Exec(r.Context(), sql, book.Title, book.AuthorID, book.Year, book.ISBN, book.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteBook handles DELETE /books/{id}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	conn, err := db.GetDB().Acquire(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	sql := "DELETE FROM books WHERE id=$1"
	_, err = conn.Exec(r.Context(), sql, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

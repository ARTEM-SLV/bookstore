package handlers

import (
	"bookstore/models"
	"bookstore/pkg/service"

	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type bookHandler struct {
	bookService service.BookService
}

var bookDB = service.NewBookService()

func NewBookHandler() *bookHandler {
	b := bookHandler{bookService: service.NewBookService()}

	return &b
}

// CreateBook handles POST /books
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := bookDB.CreateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Book created with ID: %d", id)
	w.Write([]byte(response))
}

// GetBooks handles GET /books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bookDB.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	book, err := bookDB.GetBookByID(id)
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

	err = bookDB.UpdateBook(&book)
	if err == pgx.ErrNoRows {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := "Author updated"
	w.Write([]byte(response))
}

// DeleteBook handles DELETE /books/{id}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = bookDB.DeleteBook(id)
	if err == pgx.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

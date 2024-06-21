package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"

	"bookstore/internal/logger"
	"bookstore/pkg/services"
)

type BookHandler struct {
	bookSrv services.BookService
}

func NewBookHandler(bookSrv services.BookService) *BookHandler {
	return &BookHandler{bookSrv: bookSrv}
}

// CreateBook handles POST /books
func (b *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info(fmt.Sprintf("URL: %s %s", r.URL.String(), r.Method))

	id, err := b.bookSrv.CreateBook(r.Context(), r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Book created with ID: %d", id)
	w.Write([]byte(response))
}

// GetBooks handles GET /books
func (b *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info(fmt.Sprintf("URL: %s %s", r.URL.String(), r.Method))

	books, err := b.bookSrv.GetAllBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook handles GET /books/{id}
func (b *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info(fmt.Sprintf("URL: %s %s", r.URL.String(), r.Method))

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, "Invalid book ID", http.StatusBadRequest)

		return
	}

	book, err := b.bookSrv.GetBookByID(r.Context(), id)
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
func (b *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info(fmt.Sprintf("URL: %s %s", r.URL.String(), r.Method))

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, "Invalid book ID", http.StatusBadRequest)

		return
	}

	err = b.bookSrv.UpdateBook(r.Context(), r.Body, id)
	if err == pgx.ErrNoRows {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := "Author updated"
	w.Write([]byte(response))
}

// DeleteBook handles DELETE /books/{id}
func (b *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info(fmt.Sprintf("URL: %s %s", r.URL.String(), r.Method))

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, "Invalid book ID", http.StatusBadRequest)

		return
	}

	err = b.bookSrv.DeleteBook(r.Context(), id)
	if err == pgx.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

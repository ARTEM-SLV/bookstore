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

type AuthorAndBookHandler struct {
	abSrv services.AuthorAndBookService
}

func NewAuthorAndBookHandler(abSrv services.AuthorAndBookService) *AuthorAndBookHandler {
	return &AuthorAndBookHandler{abSrv: abSrv}
}

// UpdateBookAndAuthor handles PUT /books/{book_id}/authors/{author_id}
func (ab *AuthorAndBookHandler) UpdateBookAndAuthor(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info(fmt.Sprintf("URL: %s %s", r.URL.String(), r.Method))

	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["book_id"])
	if err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, "Invalid book ID", http.StatusBadRequest)

		return
	}

	authorID, err := strconv.Atoi(vars["author_id"])
	if err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, "Invalid author ID", http.StatusBadRequest)

		return
	}

	err = ab.abSrv.UpdateBookAndAuthor(r.Context(), r.Body, bookID, authorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := "Book and Author updated"
	w.Write([]byte(response))
}

// GetAuthorAndBooks handles GET /author_with_books/{id}
func (ab AuthorAndBookHandler) GetAuthorAndBooks(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info(fmt.Sprintf("URL: %s %s", r.URL.String(), r.Method))

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		logger.Log.Error(err.Error())
		http.Error(w, "Invalid author ID", http.StatusBadRequest)

		return
	}

	author, books, err := ab.abSrv.GetAuthorAndBooks(r.Context(), id)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Author not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
	json.NewEncoder(w).Encode(books)
}

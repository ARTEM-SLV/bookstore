package handlers

import (
	"bookstore/models"
	"bookstore/pkg/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strconv"
)

type UpdateRequest struct {
	Book   models.Book        `json:"book"`
	Author models.AuthorTimeS `json:"author"`
}

// UpdateBookAndAuthor handles PUT /books/{book_id}/authors/{author_id}
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

	book := updateRequest.Book
	author := updateRequest.Author

	err = services.UpdateBookAndAuthor(&book, &author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	author, books, err := services.GetAuthorAndBooks(id)
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

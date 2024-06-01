package handlers

import (
	"bookstore/models"
	"bookstore/pkg/services"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateAuthor handles POST /authors
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.AuthorWithoutTime
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := services.CreateAuthor(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Author created with ID: %d", id)
	w.Write([]byte(response))
}

// GetAuthors handles GET /authors
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := services.GetAllAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

// GetAuthor handles GET /authors/{id}
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	author, err := services.GetAuthorByID(id)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Author not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

// UpdateAuthor handles PUT /authors/{id}
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	var author models.AuthorWithoutTime
	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	author.ID = id

	err = services.UpdateAuthor(&author)
	if err == pgx.ErrNoRows {
		http.Error(w, "Author not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := "Author updated"
	w.Write([]byte(response))
}

// DeleteAuthor handles DELETE /authors/{id}
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteAuthor(id)
	if err == pgx.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := "the author has been successfully removed along with all books by this author"
	w.Write([]byte(response))
}

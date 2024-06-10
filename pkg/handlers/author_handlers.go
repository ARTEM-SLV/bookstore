package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"

	"bookstore/pkg/services"
)

type AuthorHandler struct {
	authorSrv *services.AuthorService
}

func NewAuthorHandler(authorSrv *services.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorSrv: authorSrv}
}

// CreateAuthor handles POST /authors
func (a AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)

	id, err := a.authorSrv.CreateAuthor(r.Context(), dec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Author created with ID: %d", id)
	w.Write([]byte(response))
}

// GetAuthors handles GET /authors
func (a AuthorHandler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := a.authorSrv.GetAllAuthors(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

// GetAuthor handles GET /authors/{id}
func (a AuthorHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	author, err := a.authorSrv.GetAuthorByID(r.Context(), id)
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
func (a AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	dec := json.NewDecoder(r.Body)

	err = a.authorSrv.UpdateAuthor(r.Context(), dec, id)
	if err == pgx.ErrNoRows {
		http.Error(w, "Author not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := "Author updated"
	w.Write([]byte(response))
}

// DeleteAuthor handles DELETE /authors/{id}
func (a AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	err = a.authorSrv.DeleteAuthor(r.Context(), id)
	if err == pgx.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := "the author has been successfully removed along with all books by this author"
	w.Write([]byte(response))
}

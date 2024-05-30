package handlers

import (
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"strconv"

	"bookstore/db"
	"bookstore/pkg/models"
	"github.com/gorilla/mux"
)

// CreateAuthor handles POST /authors
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
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

	sql := `INSERT INTO authors (first_name, last_name, biography, birth_date) VALUES ($1, $2, $3, $4) RETURNING id`
	err = conn.QueryRow(r.Context(), sql, author.FirstName, author.LastName, author.Biography, author.BirthDate).Scan(&author.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

// GetAuthors handles GET /authors
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	conn, err := db.GetDB().Acquire(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(r.Context(), "SELECT id, first_name, last_name, biography, birth_date FROM authors")
	if err != nil {
		log.Println("ошибка выполнения запроса")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.ID, &author.FirstName, &author.LastName, &author.Biography, &author.BirthDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		authors = append(authors, author)
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

	conn, err := db.GetDB().Acquire(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	var author models.Author
	sql := "SELECT id, first_name, last_name, biography, birth_date FROM authors WHERE id=$1"
	err = conn.QueryRow(r.Context(), sql, id).Scan(&author.ID, &author.FirstName, &author.LastName, &author.Biography, &author.BirthDate)
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
}

// UpdateAuthor handles PUT /authors/{id}
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	var author models.Author
	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	author.ID = id

	conn, err := db.GetDB().Acquire(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	sql := "UPDATE authors SET first_name=$1, last_name=$2, biography=$3, birth_date=$4 WHERE id=$5"
	_, err = conn.Exec(r.Context(), sql, author.FirstName, author.LastName, author.Biography, author.BirthDate, author.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteAuthor handles DELETE /authors/{id}
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	conn, err := db.GetDB().Acquire(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	sql := "DELETE FROM authors WHERE id=$1"
	_, err = conn.Exec(r.Context(), sql, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

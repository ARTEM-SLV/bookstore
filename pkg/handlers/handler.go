package handlers

import (
	"net/http"
)

type BookHandler interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
	GetBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}

type AuthorHandler interface {
	CreateAuthor(w http.ResponseWriter, r *http.Request)
	GetAuthors(w http.ResponseWriter, r *http.Request)
	GetAuthor(w http.ResponseWriter, r *http.Request)
	UpdateAuthor(w http.ResponseWriter, r *http.Request)
	DeleteAuthor(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	BookHandler   BookHandler
	AuthorHandler AuthorHandler
}

func NewHandler() *Handler {
	handler := Handler{
		//BookHandler:   NewBookHandler(),
		//AuthorHandler: NewAuthorHandler(),
	}

	return &handler
}

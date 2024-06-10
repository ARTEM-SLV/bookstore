package handlers

import (
	"github.com/gorilla/mux"

	"bookstore/pkg/services"
)

func NewHandler(srv *services.Service) *mux.Router {
	b := NewBookHandler(srv.BookSrv)
	a := NewAuthorHandler(srv.AuthorSrv)
	ab := NewAuthorAndBookHandler(srv.AuthorAndBookSrv)

	r := mux.NewRouter()

	r.HandleFunc("/books", b.CreateBook).Methods("POST")
	r.HandleFunc("/books", b.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", b.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", b.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", b.DeleteBook).Methods("DELETE")

	r.HandleFunc("/authors", a.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors", a.GetAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", a.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id}", a.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", a.DeleteAuthor).Methods("DELETE")

	r.HandleFunc("/books/{book_id}/authors/{author_id}", ab.UpdateBookAndAuthor).Methods("PUT")
	r.HandleFunc("/author_with_books/{id}", ab.GetAuthorAndBooks).Methods("GET")

	return r
}

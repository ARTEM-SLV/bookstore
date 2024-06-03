package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitHandlerMux() {
	r := mux.NewRouter()

	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	r.HandleFunc("/authors", CreateAuthor).Methods("POST")
	r.HandleFunc("/authors", GetAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id}", UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", DeleteAuthor).Methods("DELETE")

	r.HandleFunc("/books/{book_id}/authors/{author_id}", UpdateBookAndAuthor).Methods("PUT")
	r.HandleFunc("/author_with_books/{id}", GetAuthorAndBooks).Methods("GET")

	http.Handle("/", r)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

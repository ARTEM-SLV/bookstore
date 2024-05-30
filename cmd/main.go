package main

import (
	"bookstore/config"
	"bookstore/db"
	"bookstore/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	StartServer()
}

func StartServer() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)

	r := mux.NewRouter()

	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", handlers.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id}", handlers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", handlers.DeleteAuthor).Methods("DELETE")

	http.Handle("/", r)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

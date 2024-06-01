package service

import (
	"bookstore/models"
	hendler "bookstore/pkg/handlers"
	"bookstore/pkg/repositories"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type BookService interface {
	CreateBook(book *models.Book) (int, error)
	GetAllBooks() ([]*models.BookAuthor, error)
	GetBookByID(id int) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id int) error
}

type AuthorService interface {
	CreateAuthor(author *models.AuthorWithoutTime) (int, error)
	GetAllAuthors() ([]*models.Author, error)
	GetAuthorByID(id int) (*models.Author, error)
	UpdateAuthor(author *models.AuthorWithoutTime) error
	DeleteAuthor(id int) error
}

type Service struct {
	bookService   BookService
	authorService AuthorService
}

var bookDB repositories.BookRepository
var authorDB repositories.AuthorRepository

func InitService() {
	//bookDB = r.BookRepository
	//authorDB = r.AuthorRepository

	StartServer()
}

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/books", hendler.CreateBook).Methods("POST")
	r.HandleFunc("/books", hendler.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", hendler.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", hendler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", hendler.DeleteBook).Methods("DELETE")

	r.HandleFunc("/authors", hendler.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors", hendler.GetAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", hendler.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id}", hendler.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", hendler.DeleteAuthor).Methods("DELETE")

	r.HandleFunc("/books/{book_id}/authors/{author_id}", hendler.UpdateBookAndAuthor).Methods("PUT")

	r.HandleFunc("/author_with_books/{id}", hendler.GetAuthorAndBooks).Methods("GET")

	http.Handle("/", r)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

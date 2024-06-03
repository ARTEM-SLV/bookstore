package repositories

import (
	"bookstore/config"
	"bookstore/models"
	"bookstore/pkg/repositories/postgre"
)

type BookRepository interface {
	CreateBook(book *models.Book) (int, error)
	GetAllBooks() ([]*models.BookAuthor, error)
	GetBookByID(id int) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id int) error
}

type AuthorRepository interface {
	CreateAuthor(author *models.AuthorTimeS) (int, error)
	GetAllAuthors() ([]*models.Author, error)
	GetAuthorByID(id int) (*models.Author, error)
	UpdateAuthor(author *models.AuthorTimeS) error
	DeleteAuthor(id int) error
}

type AuthorAndBookRepository interface {
	UpdateBookAndAuthor(book *models.Book, author *models.AuthorTimeS) error
	GetAuthorAndBooks(id int) (*models.Author, []*models.BookWithAuthor, error)
}

func InitRepositoryPG() {
	cfg := config.LoadConfig()
	postgre.InitDB(cfg)
}

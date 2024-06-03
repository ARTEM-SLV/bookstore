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
	CreateAuthor(author *models.AuthorWithoutTime) (int, error)
	GetAllAuthors() ([]*models.Author, error)
	GetAuthorByID(id int) (*models.Author, error)
	UpdateAuthor(author *models.AuthorWithoutTime) error
	DeleteAuthor(id int) error
}

func InitRepositoryPG() {
	cfg := config.LoadConfig()
	postgre.InitDB(cfg)
}

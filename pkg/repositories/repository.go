package repositories

import (
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

type Repository struct {
	BookRepository   BookRepository
	AuthorRepository AuthorRepository
}

func NewRepository() *Repository {
	postgre.InitRepository()

	repository := Repository{
		BookRepository:   postgre.NewBookRepository(),
		AuthorRepository: postgre.NewAuthorRepository(),
	}

	return &repository
}

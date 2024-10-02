package repositories

import (
	"context"

	"bookstore/config"
	"bookstore/models"
	"bookstore/pkg/repositories/postgre"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book *models.Book) (int, error)
	GetAllBooks(ctx context.Context) ([]*models.Book, error)
	GetBookByID(ctx context.Context, id int) (*models.Book, error)
	UpdateBook(ctx context.Context, book *models.Book) error
	DeleteBook(ctx context.Context, id int) error
}

type AuthorRepository interface {
	CreateAuthor(ctx context.Context, author *models.Author) (int, error)
	GetAllAuthors(ctx context.Context) ([]*models.Author, error)
	GetAuthorByID(ctx context.Context, id int) (*models.Author, error)
	UpdateAuthor(ctx context.Context, author *models.Author) error
	DeleteAuthor(ctx context.Context, id int) error
}

type AuthorAndBookRepository interface {
	UpdateBookAndAuthor(ctx context.Context, book *models.Book, author *models.Author) error
	GetAuthorAndBooks(ctx context.Context, id int) (*models.Author, []*models.Book, error)
}

type Repository struct {
	BookRepository          BookRepository
	AuthorRepository        AuthorRepository
	AuthorAndBookRepository AuthorAndBookRepository
}

func NewRepository() *Repository {
	cfg := config.LoadConfig()
	pool := postgre.InitDB(cfg)

	r := Repository{
		BookRepository:          postgre.NewBookRepository(pool),
		AuthorRepository:        postgre.NewAuthorRepository(pool),
		AuthorAndBookRepository: postgre.NewAuthorAndBookRepository(pool),
	}

	return &r
}

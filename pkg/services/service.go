package services

import (
	"context"
	"io"
	"time"

	"bookstore/models"
	"bookstore/pkg/repositories"
)

type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
	BirthDate string `json:"birth_date"`
}

type BookService interface {
	CreateBook(ctx context.Context, r io.Reader) (int, error)
	GetAllBooks(ctx context.Context) ([]*models.Book, error)
	GetBookByID(ctx context.Context, id int) (*models.Book, error)
	UpdateBook(ctx context.Context, r io.Reader, id int) error
	DeleteBook(ctx context.Context, id int) error
}

type AuthorService interface {
	CreateAuthor(ctx context.Context, r io.Reader) (int, error)
	GetAllAuthors(ctx context.Context) ([]*models.Author, error)
	GetAuthorByID(ctx context.Context, id int) (*models.Author, error)
	UpdateAuthor(ctx context.Context, r io.Reader, id int) error
	DeleteAuthor(ctx context.Context, id int) error
}

type AuthorAndBookService interface {
	UpdateBookAndAuthor(ctx context.Context, r io.Reader, bookID, authorID int) error
	GetAuthorAndBooks(ctx context.Context, id int) (*models.Author, []*models.Book, error)
}

type Service struct {
	BookSrv          *BookServicePg
	AuthorSrv        *AuthorServicePg
	AuthorAndBookSrv *AuthorAndBookServicePg
}

func NewService(rep *repositories.Repository) *Service {
	s := Service{
		BookSrv:          NewBookService(rep.BookRepository),
		AuthorSrv:        NewAuthorService(rep.AuthorRepository),
		AuthorAndBookSrv: NewAuthorAndBookService(rep.AuthorAndBookRepository),
	}

	return &s
}

func parseAuthor(receiver *models.Author, source *Author) error {
	receiver.ID = source.ID
	receiver.FirstName = source.FirstName
	receiver.LastName = source.LastName
	receiver.Biography = source.Biography

	BirthDate, err := time.Parse(time.DateOnly, source.BirthDate)
	if err != nil {
		return err
	}
	receiver.BirthDate = BirthDate

	return nil
}

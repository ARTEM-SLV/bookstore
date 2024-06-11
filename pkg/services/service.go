package services

import (
	"bookstore/models"
	"bookstore/pkg/repositories"
	"time"
)

type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
	BirthDate string `json:"birth_date"`
}

type Service struct {
	BookSrv          *BookService
	AuthorSrv        *AuthorService
	AuthorAndBookSrv *AuthorAndBookService
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

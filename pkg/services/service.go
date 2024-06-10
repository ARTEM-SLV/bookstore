package services

import (
	"bookstore/pkg/repositories"
)

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

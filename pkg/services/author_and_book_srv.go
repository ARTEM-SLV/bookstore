package services

import (
	"context"

	"bookstore/models"
	"bookstore/pkg/repositories/postgre"
)

type AuthorAndBookService struct {
	authorAndBookRep *postgre.AuthorAndBookPgRep
}

func NewAuthorAndBookService(authorAndBookRep *postgre.AuthorAndBookPgRep) *AuthorAndBookService {
	return &AuthorAndBookService{authorAndBookRep: authorAndBookRep}
}

func (a AuthorAndBookService) UpdateBookAndAuthor(ctx context.Context, book *models.Book, author *models.Author) error {
	return a.authorAndBookRep.UpdateBookAndAuthor(ctx, book, author)
}

func (a AuthorAndBookService) GetAuthorAndBooks(ctx context.Context, id int) (*models.Author, []*models.Book, error) {
	return a.authorAndBookRep.GetAuthorAndBooks(ctx, id)
}

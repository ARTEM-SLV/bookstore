package services

import (
	"bookstore/models"
	"bookstore/pkg/repositories/postgre"
	"context"
	"encoding/json"
)

type UpdateRequest struct {
	Book   models.Book `json:"book"`
	Author Author      `json:"author"`
}

type AuthorAndBookService struct {
	authorAndBookRep *postgre.AuthorAndBookPgRep
}

func NewAuthorAndBookService(authorAndBookRep *postgre.AuthorAndBookPgRep) *AuthorAndBookService {
	return &AuthorAndBookService{authorAndBookRep: authorAndBookRep}
}

func (a AuthorAndBookService) UpdateBookAndAuthor(ctx context.Context, dec *json.Decoder, bookID, authorID int) error {
	var mAuthor models.Author

	var updateRequest UpdateRequest
	err := dec.Decode(&updateRequest)
	if err != nil {
		return err
	}

	book := updateRequest.Book
	author := updateRequest.Author

	parseAuthor(&mAuthor, &author)

	return a.authorAndBookRep.UpdateBookAndAuthor(ctx, &book, &mAuthor)
}

func (a AuthorAndBookService) GetAuthorAndBooks(ctx context.Context, id int) (*models.Author, []*models.Book, error) {
	return a.authorAndBookRep.GetAuthorAndBooks(ctx, id)
}

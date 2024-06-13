package services

import (
	"context"
	"encoding/json"
	"io"

	"bookstore/models"
	"bookstore/pkg/repositories"
)

type UpdateRequest struct {
	Book   models.Book `json:"book"`
	Author Author      `json:"author"`
}

type AuthorAndBookServicePg struct {
	authorAndBookRep repositories.AuthorAndBookRepository
}

func NewAuthorAndBookService(authorAndBookRep repositories.AuthorAndBookRepository) *AuthorAndBookServicePg {
	return &AuthorAndBookServicePg{authorAndBookRep: authorAndBookRep}
}

func (a *AuthorAndBookServicePg) UpdateBookAndAuthor(ctx context.Context, r io.Reader, bookID, authorID int) error {
	var mAuthor models.Author

	var updateRequest UpdateRequest
	err := json.NewDecoder(r).Decode(&updateRequest)
	if err != nil {
		return err
	}

	book := updateRequest.Book
	author := updateRequest.Author
	book.ID = bookID
	author.ID = authorID

	parseAuthor(&mAuthor, &author)

	return a.authorAndBookRep.UpdateBookAndAuthor(ctx, &book, &mAuthor)
}

func (a *AuthorAndBookServicePg) GetAuthorAndBooks(ctx context.Context, id int) (*models.Author, []*models.Book, error) {
	return a.authorAndBookRep.GetAuthorAndBooks(ctx, id)
}

package services

import (
	"context"
	"encoding/json"
	"io"

	"bookstore/models"
	"bookstore/pkg/repositories"
)

type AuthorServicePg struct {
	authorRep repositories.AuthorRepository
}

func NewAuthorService(authorRep repositories.AuthorRepository) *AuthorServicePg {
	return &AuthorServicePg{authorRep: authorRep}
}

func (a *AuthorServicePg) CreateAuthor(ctx context.Context, r io.Reader) (int, error) {
	var mAuthor models.Author
	var author Author

	err := json.NewDecoder(r).Decode(&author)
	if err != nil {
		return 0, err
	}
	err = parseAuthor(&mAuthor, &author)
	if err != nil {
		return 0, err
	}

	return a.authorRep.CreateAuthor(ctx, &mAuthor)
}

func (a *AuthorServicePg) GetAllAuthors(ctx context.Context) ([]*models.Author, error) {
	return a.authorRep.GetAllAuthors(ctx)
}

func (a *AuthorServicePg) GetAuthorByID(ctx context.Context, id int) (*models.Author, error) {
	return a.authorRep.GetAuthorByID(ctx, id)
}

func (a *AuthorServicePg) UpdateAuthor(ctx context.Context, r io.Reader, id int) error {
	var mAuthor models.Author
	var author Author

	err := json.NewDecoder(r).Decode(&author)
	if err != nil {
		return err
	}
	err = parseAuthor(&mAuthor, &author)
	if err != nil {
		return err
	}
	mAuthor.ID = id

	return a.authorRep.UpdateAuthor(ctx, &mAuthor)
}

func (a *AuthorServicePg) DeleteAuthor(ctx context.Context, id int) error {
	return a.authorRep.DeleteAuthor(ctx, id)
}

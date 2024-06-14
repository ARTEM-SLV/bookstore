package services

import (
	"context"
	"encoding/json"
	"io"

	"bookstore/models"
	"bookstore/pkg/repositories"
)

type AuthorSrv struct {
	authorRep repositories.AuthorRepository
}

func NewAuthorService(authorRep repositories.AuthorRepository) *AuthorSrv {
	return &AuthorSrv{authorRep: authorRep}
}

func (a *AuthorSrv) CreateAuthor(ctx context.Context, r io.Reader) (int, error) {
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

func (a *AuthorSrv) GetAllAuthors(ctx context.Context) ([]*models.Author, error) {
	return a.authorRep.GetAllAuthors(ctx)
}

func (a *AuthorSrv) GetAuthorByID(ctx context.Context, id int) (*models.Author, error) {
	return a.authorRep.GetAuthorByID(ctx, id)
}

func (a *AuthorSrv) UpdateAuthor(ctx context.Context, r io.Reader, id int) error {
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

func (a *AuthorSrv) DeleteAuthor(ctx context.Context, id int) error {
	return a.authorRep.DeleteAuthor(ctx, id)
}

package services

import (
	"bookstore/models"
	"bookstore/pkg/repositories/postgre"
	"context"
	"encoding/json"
)

type AuthorService struct {
	authorRep *postgre.AuthorPgRepository
}

func NewAuthorService(authorRep *postgre.AuthorPgRepository) *AuthorService {
	return &AuthorService{authorRep: authorRep}
}

func (a AuthorService) CreateAuthor(ctx context.Context, dec *json.Decoder) (int, error) {
	var mAuthor models.Author
	var author Author

	err := dec.Decode(&author)
	if err != nil {
		return 0, err
	}
	err = parseAuthor(&mAuthor, &author)
	if err != nil {
		return 0, err
	}

	return a.authorRep.CreateAuthor(ctx, &mAuthor)
}

func (a AuthorService) GetAllAuthors(ctx context.Context) ([]*models.Author, error) {
	return a.authorRep.GetAllAuthors(ctx)
}

func (a AuthorService) GetAuthorByID(ctx context.Context, id int) (*models.Author, error) {
	return a.authorRep.GetAuthorByID(ctx, id)
}

func (a AuthorService) UpdateAuthor(ctx context.Context, dec *json.Decoder, id int) error {
	var mAuthor models.Author
	var author Author

	err := dec.Decode(&author)
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

func (a AuthorService) DeleteAuthor(ctx context.Context, id int) error {
	return a.authorRep.DeleteAuthor(ctx, id)
}

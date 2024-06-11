package services

import (
	"context"
	"encoding/json"
	"time"

	"bookstore/models"
	"bookstore/pkg/repositories/postgre"
)

type author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
	BirthDate string `json:"birth_date"`
}

type AuthorService struct {
	authorRep *postgre.AuthorPgRepository
}

func NewAuthorService(authorRep *postgre.AuthorPgRepository) *AuthorService {
	return &AuthorService{authorRep: authorRep}
}

func (a AuthorService) CreateAuthor(ctx context.Context, dec *json.Decoder) (int, error) {
	var mAuthor models.Author
	var author author

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
	var author author

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

func parseAuthor(receiver *models.Author, source *author) error {
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

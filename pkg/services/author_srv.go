package services

import (
	"context"

	"bookstore/models"
	"bookstore/pkg/repositories/postgre"
)

type AuthorService struct {
	authorRep *postgre.AuthorPgRepository
}

func NewAuthorService(authorRep *postgre.AuthorPgRepository) *AuthorService {
	return &AuthorService{authorRep: authorRep}
}

func (a AuthorService) CreateAuthor(ctx context.Context, author *models.Author) (int, error) {
	return a.authorRep.CreateAuthor(ctx, author)
}

func (a AuthorService) GetAllAuthors(ctx context.Context) ([]*models.Author, error) {
	return a.authorRep.GetAllAuthors(ctx)
}

func (a AuthorService) GetAuthorByID(ctx context.Context, id int) (*models.Author, error) {
	return a.authorRep.GetAuthorByID(ctx, id)
}

func (a AuthorService) UpdateAuthor(ctx context.Context, author *models.Author) error {
	return a.authorRep.UpdateAuthor(ctx, author)
}

func (a AuthorService) DeleteAuthor(ctx context.Context, id int) error {
	return a.authorRep.DeleteAuthor(ctx, id)
}

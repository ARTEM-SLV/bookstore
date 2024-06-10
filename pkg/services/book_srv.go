package services

import (
	"context"

	"bookstore/models"
	"bookstore/pkg/repositories/postgre"
)

type BookService struct {
	bookRep *postgre.BookPgRepository
}

func NewBookService(bookRep *postgre.BookPgRepository) *BookService {
	return &BookService{bookRep: bookRep}
}

func (b BookService) CreateBook(ctx context.Context, book *models.Book) (int, error) {
	return b.bookRep.CreateBook(ctx, book)
}

func (b BookService) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	return b.bookRep.GetAllBooks(ctx)
}

func (b BookService) GetBookByID(ctx context.Context, id int) (*models.Book, error) {
	return b.bookRep.GetBookByID(ctx, id)
}

func (b BookService) UpdateBook(ctx context.Context, book *models.Book) error {
	return b.bookRep.UpdateBook(ctx, book)
}

func (b BookService) DeleteBook(ctx context.Context, id int) error {
	return b.bookRep.DeleteBook(ctx, id)
}

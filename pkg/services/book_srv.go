package services

import (
	"context"
	"encoding/json"
	"io"

	"bookstore/models"
	"bookstore/pkg/repositories"
)

type BookSrv struct {
	bookRep repositories.BookRepository
}

func NewBookService(bookRep repositories.BookRepository) *BookSrv {
	return &BookSrv{bookRep: bookRep}
}

func (b *BookSrv) CreateBook(ctx context.Context, r io.Reader) (int, error) {
	var book models.Book
	err := json.NewDecoder(r).Decode(&book)
	if err != nil {
		return 0, err
	}

	return b.bookRep.CreateBook(ctx, &book)
}

func (b *BookSrv) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	return b.bookRep.GetAllBooks(ctx)
}

func (b *BookSrv) GetBookByID(ctx context.Context, id int) (*models.Book, error) {
	return b.bookRep.GetBookByID(ctx, id)
}

func (b *BookSrv) UpdateBook(ctx context.Context, r io.Reader, id int) error {
	var book models.Book
	err := json.NewDecoder(r).Decode(&book)
	if err != nil {
		return err
	}
	book.ID = id

	return b.bookRep.UpdateBook(ctx, &book)
}

func (b *BookSrv) DeleteBook(ctx context.Context, id int) error {
	return b.bookRep.DeleteBook(ctx, id)
}

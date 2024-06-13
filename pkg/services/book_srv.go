package services

import (
	"context"
	"encoding/json"
	"io"

	"bookstore/models"
	"bookstore/pkg/repositories"
	"bookstore/pkg/repositories/postgre"
)

type BookServicePg struct {
	bookRep repositories.BookRepository
}

func NewBookService(bookRep *postgre.BookRepositoryPg) *BookServicePg {
	return &BookServicePg{bookRep: bookRep}
}

func (b *BookServicePg) CreateBook(ctx context.Context, r io.Reader) (int, error) {
	var book models.Book
	err := json.NewDecoder(r).Decode(&book)
	if err != nil {
		return 0, err
	}

	return b.bookRep.CreateBook(ctx, &book)
}

func (b *BookServicePg) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	return b.bookRep.GetAllBooks(ctx)
}

func (b *BookServicePg) GetBookByID(ctx context.Context, id int) (*models.Book, error) {
	return b.bookRep.GetBookByID(ctx, id)
}

func (b *BookServicePg) UpdateBook(ctx context.Context, r io.Reader, id int) error {
	var book models.Book
	err := json.NewDecoder(r).Decode(&book)
	if err != nil {
		return err
	}
	book.ID = id

	return b.bookRep.UpdateBook(ctx, &book)
}

func (b *BookServicePg) DeleteBook(ctx context.Context, id int) error {
	return b.bookRep.DeleteBook(ctx, id)
}

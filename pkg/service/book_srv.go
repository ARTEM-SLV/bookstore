package service

import (
	"bookstore/models"
)

type bookService struct {
}

func NewBookService() *bookService {
	return &bookService{}
}

func (b bookService) CreateBook(book *models.Book) (int, error) {
	return bookDB.CreateBook(book)
}

func (b bookService) GetAllBooks() ([]*models.BookAuthor, error) {
	return bookDB.GetAllBooks()
}

func (b bookService) GetBookByID(id int) (*models.Book, error) {
	return bookDB.GetBookByID(id)
}

func (b bookService) UpdateBook(book *models.Book) error {
	return bookDB.UpdateBook(book)
}

func (b bookService) DeleteBook(id int) error {
	return bookDB.DeleteBook(id)
}

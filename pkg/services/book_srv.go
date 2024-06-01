package services

import (
	"bookstore/models"
	"bookstore/pkg/repositories/postgre"
)

var bookRep = postgre.NewBookRepository()

func CreateBook(book *models.Book) (int, error) {
	return bookRep.CreateBook(book)
}

func GetAllBooks() ([]*models.BookAuthor, error) {
	return bookRep.GetAllBooks()
}

func GetBookByID(id int) (*models.Book, error) {
	return bookRep.GetBookByID(id)
}

func UpdateBook(book *models.Book) error {
	return bookRep.UpdateBook(book)
}

func DeleteBook(id int) error {
	return bookRep.DeleteBook(id)
}

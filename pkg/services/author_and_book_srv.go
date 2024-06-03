package services

import "bookstore/models"

func UpdateBookAndAuthor(book *models.Book, author *models.AuthorTimeS) error {
	return authorAndBookRep.UpdateBookAndAuthor(book, author)
}

func GetAuthorAndBooks(id int) (*models.Author, []*models.BookWithAuthor, error) {
	return authorAndBookRep.GetAuthorAndBooks(id)
}

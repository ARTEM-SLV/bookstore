package service

import "bookstore/models"

type authorService struct{}

func NewAuthorService() *authorService {
	return new(authorService)
}

func CreateAuthor(author *models.AuthorWithoutTime) (int, error) {
	return authorDB.CreateAuthor(author)
}
func GetAllAuthors() ([]*models.Author, error) {
	return authorDB.GetAllAuthors()
}
func GetAuthorByID(id int) (*models.Author, error) {
	return authorDB.GetAuthorByID(id)
}
func UpdateAuthor(author *models.AuthorWithoutTime) error {
	return authorDB.UpdateAuthor(author)
}
func DeleteAuthor(id int) error {
	return authorDB.DeleteAuthor(id)
}

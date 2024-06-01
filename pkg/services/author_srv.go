package services

import (
	"bookstore/models"
	"bookstore/pkg/repositories/postgre"
)

var authorRep = postgre.NewAuthorRepository()

func CreateAuthor(author *models.AuthorWithoutTime) (int, error) {
	return authorRep.CreateAuthor(author)
}
func GetAllAuthors() ([]*models.Author, error) {
	return authorRep.GetAllAuthors()
}
func GetAuthorByID(id int) (*models.Author, error) {
	return authorRep.GetAuthorByID(id)
}
func UpdateAuthor(author *models.AuthorWithoutTime) error {
	return authorRep.UpdateAuthor(author)
}
func DeleteAuthor(id int) error {
	return authorRep.DeleteAuthor(id)
}

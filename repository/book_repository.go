package repository

import (
	"enigmacamp.com/tddbook/model"
	"github.com/google/uuid"
)

type IBookRepository interface {
	Insert(b model.Book) *model.Book
}
type BookRepository struct {
	db []model.Book
}

func (repo *BookRepository) Insert(b model.Book) *model.Book {
	uid, _ := uuid.NewUUID()
	b.Id = uid.String()
	repo.db = append(repo.db, b)
	return &b
}
func NewBookRepository() IBookRepository {
	return &BookRepository{}
}

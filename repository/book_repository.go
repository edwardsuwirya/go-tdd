package repository

import (
	"enigmacamp.com/tddbook/model"
	"github.com/google/uuid"
)

type IBookRepository interface {
	Insert(b model.Book) *model.Book
}
type BookRepository struct {
	db *InMemDB
}

func (repo *BookRepository) Insert(b model.Book) *model.Book {
	uid, _ := uuid.NewUUID()
	b.Id = uid.String()
	repo.db.Insert(b)
	return b
}
func NewBookRepository(db *InMemDB) BookRepository {
	return BookRepository{
		db: db,
	}
}

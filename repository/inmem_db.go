package repository

import "enigmacamp.com/tddbook/model"

type InMemDB struct {
	Books []model.Book
}

func (db *InMemDB) Insert(b model.Book) model.Book {
	db.Books = append(db.Books, b)
	return b
}
func (db *InMemDB) GetAll() []model.Book {
	return db.Books
}
func NewInMemDB() *InMemDB {
	return &InMemDB{}
}

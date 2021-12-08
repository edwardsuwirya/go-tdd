package repository

import (
	"enigmacamp.com/tddbook/model"
	"testing"
)

var dummyBook = model.Book{
	Name:      "Dummy book 1",
	Author:    "Joni",
	Pages:     111,
	Publisher: "Erlangga",
}

func TestCanInsertBook(t *testing.T) {
	t.Run("Insert Dummy Book to repository", func(t *testing.T) {
		bookRepository := NewBookRepository()
		got := bookRepository.Insert(dummyBook)
		expected := model.Book{
			Id:        "123",
			Name:      "Dummy book 1",
			Author:    "Joni",
			Pages:     111,
			Publisher: "Erlangga",
		}
		if got.Id == "" {
			t.Errorf("Got: %v, have no ID", got)
		}
		if got.Name != expected.Name {
			t.Errorf("Got: %v, Expected: %v", got.Name, expected.Name)
		}
		if len(bookRepository.(*BookRepository).db) != 1 {
			t.Errorf("Insert book failed")
		}
	},
	)
}

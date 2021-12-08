package usecase

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

type MockRepository struct {
}

func (repo *MockRepository) Insert(b model.Book) *model.Book {
	b.Id = "123"
	return &b
}

func TestCanRegisterBook(t *testing.T) {
	t.Run("Register Dummy Book 1", func(t *testing.T) {
		mockRepo := new(MockRepository)
		registerBookUseCase := NewRegisterBookUseCase(mockRepo)
		got := registerBookUseCase.NewRegistration(dummyBook)
		expected := mockRepo.Insert(dummyBook)

		if got.Id != expected.Id {
			t.Errorf("Got: %v, have no ID", got)
		}
		if got.Name != expected.Name {
			t.Errorf("Got: %v, Expected: %v", got.Name, expected.Name)
		}
	})
}

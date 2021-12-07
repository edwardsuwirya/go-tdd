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

func TestCanRegisterBook(t *testing.T) {
	t.Run("Register Dummy Book 1", func(t *testing.T) {
		registerBookUseCase := NewRegisterBookUseCase()
		got := registerBookUseCase.NewRegistration(dummyBook)
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
	})
}

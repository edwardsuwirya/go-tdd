package usecase

import (
	"testing"
)

var dummyBook = Book{
	Name:      "Dummy book 1",
	Author:    "Joni",
	Pages:     111,
	Publisher: "Erlangga",
}

func TestCanRegisterBook(t *testing.T) {
	t.Run("Register Dummy Book 1", func(t *testing.T) {
		registerBookUseCase := NewRegisterBookUseCase()
		got := registerBookUseCase.NewRegistration(dummyBook)
		expected := Book{
			Id:        "123",
			Name:      "Dummy name 1",
			Author:    "Joni",
			Pages:     111,
			Publisher: "Erlangga",
		}

		if got.Id == "" {
			t.Errorf("Got: %v, have no ID", got)
		}
		if got.Name != expected.Name {
			t.Errorf("Got: %v, Expected: %v", got, expected)
		}
	})
}

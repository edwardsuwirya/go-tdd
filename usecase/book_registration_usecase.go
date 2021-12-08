package usecase

import (
	"enigmacamp.com/tddbook/model"
	"enigmacamp.com/tddbook/repository"
)

type IRegisterBookUseCase interface {
	NewRegistration(b model.Book) *model.Book
}
type RegisterBookUseCase struct {
	repo repository.IBookRepository
}

func (uc *RegisterBookUseCase) NewRegistration(b model.Book) *model.Book {
	bookRegistered := uc.repo.Insert(b)
	return bookRegistered
}
func NewRegisterBookUseCase(repo repository.IBookRepository) IRegisterBookUseCase {
	return &RegisterBookUseCase{
		repo: repo,
	}
}

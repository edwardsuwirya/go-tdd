package usecase

import (
	"enigmacamp.com/tddbook/model"
	"enigmacamp.com/tddbook/repository"
	"github.com/google/uuid"
)

type IRegisterBookUseCase interface {
	NewRegistration(b model.Book) model.Book
}
type RegisterBookUseCase struct {
	repo repository.IBookRepository
}

func (uc *RegisterBookUseCase) NewRegistration(b model.Book) model.Book {
	uid, _ := uuid.NewUUID()
	b.Id = uid.String()
	return b
}
func NewRegisterBookUseCase(repo repository.IBookRepository) IRegisterBookUseCase {
	return &RegisterBookUseCase{
		repo: repo,
	}
}

package usecase

import (
	"enigmacamp.com/tddbook/model"
	"github.com/google/uuid"
)

type IRegisterBookUseCase interface {
	NewRegistration(b model.Book) model.Book
}
type RegisterBookUseCase struct {
}

func (uc *RegisterBookUseCase) NewRegistration(b model.Book) model.Book {
	uid, _ := uuid.NewUUID()
	b.Id = uid.String()
	return b
}
func NewRegisterBookUseCase() IRegisterBookUseCase {
	return &RegisterBookUseCase{}
}

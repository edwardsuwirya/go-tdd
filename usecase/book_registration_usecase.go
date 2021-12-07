package usecase

type RegisterBookUseCase struct {
}

func (uc *RegisterBookUseCase) NewRegistration(b Book) Book {
	return Book{
		Id:        "123",
		Name:      "Dummy name 1",
		Author:    "Joni",
		Pages:     111,
		Publisher: "Erlangga",
	}
}
func NewRegisterBookUseCase() RegisterBookUseCase {
	return RegisterBookUseCase{}
}

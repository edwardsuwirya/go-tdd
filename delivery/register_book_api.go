package delivery

import (
	"enigmacamp.com/tddbook/model"
	"enigmacamp.com/tddbook/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BookApi struct {
	usecase usecase.IRegisterBookUseCase
}

func (api *BookApi) InitRouter(router *gin.Engine) {
	router.POST("", api.RegisterBook)
}

func (api *BookApi) RegisterBook(c *gin.Context) {
	var book model.Book
	err := c.BindJSON(&book)
	if err != nil {
		fmt.Println(err)
	}
	bookRegistered := api.usecase.NewRegistration(book)
	c.JSON(200, bookRegistered)
}
func NewBookApi(useCase usecase.IRegisterBookUseCase) *BookApi {
	return &BookApi{
		usecase: useCase,
	}
}

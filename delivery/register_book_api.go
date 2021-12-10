package delivery

import "github.com/gin-gonic/gin"

type BookApi struct {
}

func (api *BookApi) InitRouter(router *gin.Engine) {
	router.POST("", api.RegisterBook)
}

func (api *BookApi) RegisterBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "",
	})
}
func NewBookApi() *BookApi {
	return &BookApi{}
}

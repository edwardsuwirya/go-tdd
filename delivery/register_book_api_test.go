package delivery

import (
	"bytes"
	"encoding/json"
	"enigmacamp.com/tddbook/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

var routerTest = gin.Default()
var dummyBook = model.Book{
	Name:      "Dummy book 1",
	Author:    "Joni",
	Pages:     111,
	Publisher: "Erlangga",
}

type RegisterBookUseCaseMock struct {
}

func (uc *RegisterBookUseCaseMock) NewRegistration(b model.Book) model.Book {
	return b
}

func TestBookApi(t *testing.T) {
	t.Run("Call Register Book API", func(t *testing.T) {
		bookApi := NewBookApi(&RegisterBookUseCaseMock{})
		bookApi.InitRouter(routerTest)
		handler := bookApi.RegisterBook
		routerTest.POST("/book", handler)
		rr := httptest.NewRecorder()

		reqBody, _ := json.Marshal(dummyBook)
		request, _ := http.NewRequest(http.MethodPost, "/book", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")
		routerTest.ServeHTTP(rr, request)

		responseCodeGot := rr.Code
		responseCodeExpected := 200

		if responseCodeGot != responseCodeExpected {
			t.Errorf("Response Code Got :%v, expected: %v", responseCodeGot, responseCodeExpected)
		}
		responseBodyGot := rr.Body.String()
		var actualResponseJson model.Book
		json.Unmarshal([]byte(responseBodyGot), &actualResponseJson)
		responseBodyExpected := dummyBook
		if actualResponseJson.Name != responseBodyExpected.Name {
			t.Errorf("Response Body Got:%v, expected: %v", responseBodyGot, responseBodyExpected)
		}
	})
}

package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

var routerTest = gin.Default()

func TestBookApi(t *testing.T) {
	t.Run("Call Register Book API", func(t *testing.T) {
		bookApi := NewBookApi()
		bookApi.InitRouter(routerTest)
		handler := bookApi.RegisterBook
		routerTest.POST("/book", handler)
		rr := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, "/book", nil)
		routerTest.ServeHTTP(rr, request)

		responseCodeGot := rr.Code
		responseCodeExpected := 200

		if responseCodeGot != responseCodeExpected {
			t.Errorf("Response Code Got :%v, expected: %v", responseCodeGot, responseCodeExpected)
		}
		responseBodyGot := rr.Body.String()
		responseBodyExpected := "{\"message\":\"\"}"
		if responseBodyGot != responseBodyExpected {
			t.Errorf("Response Body Got:%v, expected: %v", responseBodyGot, responseBodyExpected)
		}
	})
}

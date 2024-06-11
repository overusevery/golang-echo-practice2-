package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase"
)

func TestCustomerHandler_GetCustomer(t *testing.T) {
	h := &CustomerHandler{
		GetCustomerUseCase: usecase.GetCustomerUseCase{},
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetPath("/customer/:id")
	c.SetParamNames("id")
	c.SetParamValues("12")

	h.GetCustomer(c)
	assertResponseBody(t, res.Body.String(), "{id:122}")
	fmt.Println("res.Body.String()")
	fmt.Println(res.Result().StatusCode)
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

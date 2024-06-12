package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCustomerHandler_GetCustomer(t *testing.T) {
	h := &CustomerHandler{
		GetCustomerUseCase: usecase.GetCustomerUseCase{},
	}
	e := echo.New()
	h.RegisterRouter(e)
	req := httptest.NewRequest(http.MethodGet, "/customer/12", nil)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	expectedJson, err := os.ReadFile("../../fixture/customer.json")
	if err != nil {
		panic(err)
	}

	assert.JSONEq(t, string(expectedJson), res.Body.String())
}

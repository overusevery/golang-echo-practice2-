package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase"
	mock_repository "github.com/overusevery/golang-echo-practice2/src/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCustomerHandler_GetCustomer(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t) //t *testing.T)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	m.EXPECT().GetCustomer(gomock.Eq(12)).Return(entity.Customer{ID: 12})
	h := NewCustomrHandler(*usecase.NewGetCustomerUseCase(m))
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

package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

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
	m.EXPECT().GetCustomer(gomock.Eq(1)).Return(entity.Customer{
		ID:            1,
		Name:          "山田 太郎",
		Address:       "東京都練馬区豊玉北2-13-1",
		ZIP:           "176-0013",
		Phone:         "03-1234-5678",
		MarketSegment: "個人",
		Nation:        "日本",
		Birthdate:     time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	h := NewCustomrHandler(*usecase.NewGetCustomerUseCase(m))
	h.RegisterRouter(e)
	req := httptest.NewRequest(http.MethodGet, "/customer/1", nil)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	expectedJson, err := os.ReadFile("../../fixture/customer.json")
	if err != nil {
		panic(err)
	}

	assert.JSONEq(t, string(expectedJson), res.Body.String())
}

func TestCustomerHandler_CreateCustomer(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t) //t *testing.T)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	m.EXPECT().CreateCustomer(gomock.Eq(entity.Customer{
		ID:            1,
		Name:          "山田 太郎",
		Address:       "東京都練馬区豊玉北2-13-1",
		ZIP:           "176-0013",
		Phone:         "03-1234-5678",
		MarketSegment: "個人",
		Nation:        "日本",
		Birthdate:     time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	})).Return(nil)
	h := NewCreateCustomerHandler(*usecase.NewCreateCustomerUseCase(m))
	h.RegisterRouter(e)
	requestJson, err := os.ReadFile("../../fixture/customer.json")
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/customer", bytes.NewReader(requestJson))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	expectedJson, err := os.ReadFile("../../fixture/create_customer_success.json")
	if err != nil {
		panic(err)
	}

	assert.JSONEq(t, string(expectedJson), res.Body.String())
}

package handler

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	mock_repository "github.com/overusevery/golang-echo-practice2/src/repository/mock"
	"github.com/overusevery/golang-echo-practice2/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCustomerHandler_GetCustomer(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	m.EXPECT().GetCustomer(context.Background(), gomock.Eq(1)).Return(&entity.Customer{
		ID:            1,
		Name:          "山田 太郎",
		Address:       "東京都練馬区豊玉北2-13-1",
		ZIP:           "176-0013",
		Phone:         "03-1234-5678",
		MarketSegment: "個人",
		Nation:        "日本",
		Birthdate:     time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	}, nil)
	h := NewGetCustomrHandler(customerusecase.NewGetCustomerUseCase(m))
	h.RegisterRouter(e)

	res := testutil.GET(e, "/customer/1")

	expectedJson, err := os.ReadFile("../../fixture/get_customer_response.json")
	if err != nil {
		panic(err)
	}

	assert.JSONEq(t, string(expectedJson), res.Body.String())
}

func TestCustomerHandler_GetCustomer_NotFound(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	m.EXPECT().GetCustomer(context.Background(), gomock.Eq(1)).Return(nil, repository.ErrCustomerNotFound)
	h := NewGetCustomrHandler(customerusecase.NewGetCustomerUseCase(m))
	h.RegisterRouter(e)

	res := testutil.GET(e, "/customer/1")

	assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
}

func TestCustomerHandler_When_Unexpected_Error_Happened_InternalServerError_Should_be_Returned(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	m.EXPECT().GetCustomer(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))
	h := NewGetCustomrHandler(customerusecase.NewGetCustomerUseCase(m))
	h.RegisterRouter(e)

	res := testutil.GET(e, "/customer/1")

	assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
}

func TestCustomerHandler_Invalid_Query_Should_Return_Bad_Request(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	h := NewGetCustomrHandler(customerusecase.NewGetCustomerUseCase(m))
	h.RegisterRouter(e)

	res := testutil.GET(e, "/customer/not_number")

	assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
}

func TestCustomerHandler_CreateCustomer(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	m.EXPECT().CreateCustomer(context.Background(), gomock.Eq(entity.Customer{
		ID:            1,
		Name:          "山田 太郎",
		Address:       "東京都練馬区豊玉北2-13-1",
		ZIP:           "176-0013",
		Phone:         "03-1234-5678",
		MarketSegment: "個人",
		Nation:        "日本",
		Birthdate:     time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	})).Return(&entity.Customer{
		ID:            1,
		Name:          "山田 太郎",
		Address:       "東京都練馬区豊玉北2-13-1",
		ZIP:           "176-0013",
		Phone:         "03-1234-5678",
		MarketSegment: "個人",
		Nation:        "日本",
		Birthdate:     time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	}, nil)
	h := NewCreateCustomerHandler(customerusecase.NewCreateCustomerUseCase(m))
	h.RegisterRouter(e)
	res := testutil.Post(e, "/customer", "../../fixture/create_customer_request.json")
	expectedJson, err := os.ReadFile("../../fixture/create_customer_response.json")
	if err != nil {
		panic(err)
	}

	assert.JSONEq(t, string(expectedJson), res.Body.String())
}

func TestCustomerHandler_CreateCustomer_Bad_Request(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	h := NewCreateCustomerHandler(customerusecase.NewCreateCustomerUseCase(m))
	h.RegisterRouter(e)
	res := testutil.Post(e, "/customer", "../../fixture/create_customer_request_invalid.json")
	assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
}

func TestCustomerHandler_CreateCustomer_When_Unexpected_Error_Happened_InternalServerError_Should_be_Returned(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	m.EXPECT().CreateCustomer(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))
	h := NewCreateCustomerHandler(customerusecase.NewCreateCustomerUseCase(m))
	h.RegisterRouter(e)
	res := testutil.Post(e, "/customer", "../../fixture/create_customer_request.json")
	assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
}

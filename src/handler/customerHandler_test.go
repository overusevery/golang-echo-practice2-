package handler

import (
	"context"
	"errors"
	"net/http"
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

func Test_GetCustomer(t *testing.T) {
	setupGetCustomerHandlerWithMock(t, func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
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

		res := testutil.GET(e, "/customer/1")

		testutil.AssertResBodyIsEquWithJson(t, res, "../../fixture/get_customer_response.json")
	})
}

func Test_GetCustomer_NotFound(t *testing.T) {
	setupGetCustomerHandlerWithMock(t, func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
		m.EXPECT().GetCustomer(context.Background(), gomock.Eq(1)).Return(nil, repository.ErrCustomerNotFound)

		res := testutil.GET(e, "/customer/1")

		assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	})
}

func Test_GetCustomer_InternalServerError(t *testing.T) {
	setupGetCustomerHandlerWithMock(t, func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
		m.EXPECT().GetCustomer(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))

		res := testutil.GET(e, "/customer/1")

		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
}

func Test_GetCustomer_BadRequest(t *testing.T) {
	setupGetCustomerHandlerWithMock(t, func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {

		res := testutil.GET(e, "/customer/not_number")

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})
}

func Test_CreateCustomer(t *testing.T) {
	setupCreateCustomerHandlerWithMock(t,
		func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
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

			res := testutil.Post(e, "/customer", "../../fixture/create_customer_request.json")

			testutil.AssertResBodyIsEquWithJson(t, res, "../../fixture/create_customer_response.json")
		})
}

func Test_CreateCustomer_Bad_Request(t *testing.T) {
	setupCreateCustomerHandlerWithMock(t,
		func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {

			res := testutil.Post(e, "/customer", "../../fixture/create_customer_request_invalid.json")

			assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
		})
}

func Test_CreateCustomer_InternalServerError(t *testing.T) {
	setupCreateCustomerHandlerWithMock(t,
		func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
			m.EXPECT().CreateCustomer(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))

			res := testutil.Post(e, "/customer", "../../fixture/create_customer_request.json")

			assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
		})
}

func setupGetCustomerHandlerWithMock(t *testing.T, testFun func(m *mock_repository.MockCustomerRepository, e *echo.Echo)) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	h := NewGetCustomrHandler(customerusecase.NewGetCustomerUseCase(m))
	h.RegisterRouter(e)

	testFun(m, e)
}

func setupCreateCustomerHandlerWithMock(t *testing.T, testFun func(m *mock_repository.MockCustomerRepository, e *echo.Echo)) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	h := NewCreateCustomerHandler(customerusecase.NewCreateCustomerUseCase(m))
	h.RegisterRouter(e)

	testFun(m, e)
}

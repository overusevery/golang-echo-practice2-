package customerhandler

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	"github.com/overusevery/golang-echo-practice2/src/handler/customemiddleware"
	mock_repository "github.com/overusevery/golang-echo-practice2/src/repository/mock"
	"github.com/overusevery/golang-echo-practice2/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupCreateCustomerHandlerWithMock(t,
			func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
				m.EXPECT().CreateCustomer(gomock.Any(), CustomerMatcherButID{expected: *forceNewCustomer(
					"0",
					"山田 太郎",
					"東京都練馬区豊玉北2-13-1",
					"176-0013",
					"03-1234-5678",
					"個人",
					"日本",
					time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
					1,
				)}).Return(forceNewCustomer(
					"1",
					"山田 太郎",
					"東京都練馬区豊玉北2-13-1",
					"176-0013",
					"03-1234-5678",
					"個人",
					"日本",
					time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
					1,
				), nil)

				res := testutil.Post(e, "/customer", "../../../fixture/create_customer_request.json")

				testutil.AssertResBodyIsEquWithJson(t, res, "../../../fixture/create_customer_response.json")
			})
	})
	t.Run("bad request(request json validation)", func(t *testing.T) {
		setupCreateCustomerHandlerWithMock(t,
			func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {

				res := testutil.Post(e, "/customer", "../../../fixture/create_customer_request_invalid.json")

				assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
				testutil.AssertResBodyIsEquWithJson(t, res, "../../../fixture/create_customer_response_error_message_ex_ERRID00004.json")
			})
	})
	t.Run("bad request(domain model validation)", func(t *testing.T) {
		setupCreateCustomerHandlerWithMock(t,
			func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {

				res := testutil.Post(e, "/customer", "../../../fixture/create_customer_request_invalid_too_old_birthdate.json")

				assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
			})
	})
	t.Run("domain model validation error should return error id list", func(t *testing.T) {
		t.Run("single error", func(t *testing.T) {
			testCases := []struct {
				name             string
				inputJsonPath    string
				expectedJsonPath string
			}{
				{
					name:             "ERRID00001",
					inputJsonPath:    "../../../fixture/create_customer_request_invalid_too_old_birthdate.json",
					expectedJsonPath: "../../../fixture/create_customer_response_single_error_message_ex_ERRID00001.json",
				},
				{
					name:             "ERRID00002",
					inputJsonPath:    "../../../fixture/create_customer_request_invalid_future_birthdate.json",
					expectedJsonPath: "../../../fixture/create_customer_response_single_error_message_ex_ERRID00002.json",
				},
				{
					name:             "ERRID00003",
					inputJsonPath:    "../../../fixture/create_customer_request_invalid_nation.json",
					expectedJsonPath: "../../../fixture/create_customer_response_single_error_message_ex_ERRID00003.json",
				},
			}
			for _, c := range testCases {
				t.Run(c.name, func(t *testing.T) {
					setupCreateCustomerHandlerWithMock(t,
						func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {

							res := testutil.Post(e, "/customer", c.inputJsonPath)

							assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
							testutil.AssertResBodyIsEquWithJson(t, res, c.expectedJsonPath)
						})

				})
			}
		})
		t.Run("multiple error", func(t *testing.T) {
			setupCreateCustomerHandlerWithMock(t,
				func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {

					res := testutil.Post(e, "/customer", "../../../fixture/create_customer_request_invalid_multiple_error.json")

					assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
					testutil.AssertResBodyIsEquWithJson(t, res, "../../../fixture/create_customer_response_multiple_error_message_ex_ERRID00002_3.json")
				})
		})
	})
	t.Run("internal server error", func(t *testing.T) {
		setupCreateCustomerHandlerWithMock(t,
			func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
				m.EXPECT().CreateCustomer(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))

				res := testutil.Post(e, "/customer", "../../../fixture/create_customer_request.json")

				assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
			})
	})
}

func setupCreateCustomerHandlerWithMock(t *testing.T, testFun func(m *mock_repository.MockCustomerRepository, e *echo.Echo)) {
	e := echo.New()
	e.Use(customemiddleware.ParseAuthorizationToken)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	h := NewCreateCustomerHandler(customerusecase.NewCreateCustomerUseCase(m))
	h.RegisterRouter(e)

	testFun(m, e)
}

func forceNewCustomer(id string, name string, address string, zip string, phone string, marketSegment string, nation string, birthdate time.Time, version int) *entity.Customer {
	c, err := entity.NewCustomer(
		id,
		name,
		address,
		zip,
		phone,
		marketSegment,
		nation,
		birthdate,
		version,
	)
	if err != nil {
		panic(fmt.Errorf("failed to create customer in test code:%w", err))
	}
	return c
}

type CustomerMatcherButID struct {
	expected entity.Customer
}

func (m CustomerMatcherButID) Matches(x interface{}) bool {
	actual, ok := x.(entity.Customer)
	if !ok {
		return false
	}
	actual.ID = m.expected.ID
	return reflect.DeepEqual(actual, m.expected)
}

func (m CustomerMatcherButID) String() string {
	return "matches ignoring ID"
}

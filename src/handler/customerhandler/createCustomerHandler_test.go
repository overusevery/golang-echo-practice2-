package customerhandler

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	mock_repository "github.com/overusevery/golang-echo-practice2/src/repository/mock"
	"github.com/overusevery/golang-echo-practice2/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupCreateCustomerHandlerWithMock(t,
			func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
				m.EXPECT().CreateCustomer(context.Background(), gomock.Eq(entity.Customer{
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

				res := testutil.Post(e, "/customer", "../../../fixture/create_customer_request.json")

				testutil.AssertResBodyIsEquWithJson(t, res, "../../../fixture/create_customer_response.json")
			})
	})
	t.Run("bad request", func(t *testing.T) {
		setupCreateCustomerHandlerWithMock(t,
			func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {

				res := testutil.Post(e, "/customer", "../../../fixture/create_customer_request_invalid.json")

				assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)
	h := NewCreateCustomerHandler(customerusecase.NewCreateCustomerUseCase(m))
	h.RegisterRouter(e)

	testFun(m, e)
}

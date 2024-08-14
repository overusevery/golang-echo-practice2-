package customerhandler

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	"github.com/overusevery/golang-echo-practice2/src/domain/value"
	mock_repository "github.com/overusevery/golang-echo-practice2/src/repository/mock"
	"github.com/overusevery/golang-echo-practice2/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupGetCustomerHandlerWithMock(t, func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
			m.EXPECT().GetCustomer(context.Background(), gomock.Eq(value.NewID("1"))).Return(forceNewCustomer(
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

			res := testutil.GET(e, "/customer/1")

			testutil.AssertResBodyIsEquWithJson(t, res, "../../../fixture/get_customer_response.json")
		})
	})
	t.Run("not found", func(t *testing.T) {
		setupGetCustomerHandlerWithMock(t, func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
			m.EXPECT().GetCustomer(context.Background(), gomock.Eq(value.NewID("1"))).Return(nil, repository.ErrCustomerNotFound)

			res := testutil.GET(e, "/customer/1")

			assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
		})
	})
	t.Run("bad request", func(t *testing.T) {
		t.Skip() //no such case at now
		// setupGetCustomerHandlerWithMock(t, func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {

		// 	res := testutil.GET(e, "/customer/not_number")

		// 	assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
		// })
	})
	t.Run("internal server error", func(t *testing.T) {
		setupGetCustomerHandlerWithMock(t, func(m *mock_repository.MockCustomerRepository, e *echo.Echo) {
			m.EXPECT().GetCustomer(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))

			res := testutil.GET(e, "/customer/1")

			assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
		})
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

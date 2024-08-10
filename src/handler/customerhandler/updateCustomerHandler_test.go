package customerhandler

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	mock_repository "github.com/overusevery/golang-echo-practice2/src/repository/mock"
	"github.com/overusevery/golang-echo-practice2/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateCustomerHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		e, m, close := setUpdateCustomerMock(t)
		defer close()
		m.EXPECT().UpdateCustomer(gomock.Any(), gomock.Eq(*forceNewCustomer(
			"1",
			"new name",
			"new address",
			"xxx-xxx-xxx",
			"11-111-111",
			"company",
			"JP",
			time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		))).Return(forceNewCustomer(
			"1",
			"new name",
			"new address",
			"xxx-xxx-xxx",
			"11-111-111",
			"company",
			"JP",
			time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		), nil)
		res := testutil.PUT(e, "/customer/1", "../../../fixture/put_customer_request.json")
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		testutil.AssertResBodyIsEquWithJson(t, res, "../../../fixture/put_customer_response.json")
	})
	t.Run("not found", func(t *testing.T) {
		e, m, close := setUpdateCustomerMock(t)
		defer close()
		m.EXPECT().UpdateCustomer(gomock.Any(), gomock.Any()).Return(forceNewCustomer(
			"1",
			"new name",
			"new address",
			"xxx-xxx-xxx",
			"11-111-111",
			"company",
			"JP",
			time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		), repository.ErrCustomerNotFound)
		res := testutil.PUT(e, "/customer/notexist", "../../../fixture/put_customer_request.json")
		assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	})
	t.Run("internal server error", func(t *testing.T) {
		e, m, close := setUpdateCustomerMock(t)
		defer close()
		m.EXPECT().UpdateCustomer(gomock.Any(), gomock.Any()).Return(forceNewCustomer(
			"1",
			"new name",
			"new address",
			"xxx-xxx-xxx",
			"11-111-111",
			"company",
			"JP",
			time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		), errors.New("some error"))
		res := testutil.PUT(e, "/customer/1", "../../../fixture/put_customer_request.json")
		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
}

func setUpdateCustomerMock(t *testing.T) (*echo.Echo, *mock_repository.MockCustomerRepository, func()) {
	e := echo.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_repository.NewMockCustomerRepository(ctrl)

	h := &UpdateCustomerHandler{
		UpdateCustomerUseCase: &customerusecase.UpdateCustomerUseCase{Repository: m},
	}
	h.RegisterRouter(e)
	return e, m, ctrl.Finish
}

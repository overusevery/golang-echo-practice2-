package customerhandler

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
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
		m.EXPECT().UpdateCustomer(gomock.Any(), gomock.Any())
		res := testutil.PUT(e, "/customer/1", "../../../fixture/put_customer_request.json")
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("fake:internal server error", func(t *testing.T) {
		e, m, close := setUpdateCustomerMock(t)
		defer close()
		m.EXPECT().UpdateCustomer(gomock.Any(), gomock.Any())
		res := testutil.PUT(e, "/customer/2", "../../../fixture/put_customer_request.json")
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

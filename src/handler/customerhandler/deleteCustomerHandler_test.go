package customerhandler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	mock_repository "github.com/overusevery/golang-echo-practice2/src/repository/mock"
	"github.com/overusevery/golang-echo-practice2/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteCustomerHandler_DeleteCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		e, close, m := setupMock(t)
		defer close()

		m.EXPECT().DeleteCustomer(gomock.Any(), "1")
		res := testutil.DELETE(e, "/customer/1")
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("not found", func(t *testing.T) {
		e, close, m := setupMock(t)
		defer close()

		m.EXPECT().DeleteCustomer(gomock.Any(), "1").Return(repository.ErrCustomerNotFound)
		res := testutil.DELETE(e, "/customer/1")
		assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	})
	t.Run("internal server error", func(t *testing.T) {
		e, close, m := setupMock(t)
		defer close()

		m.EXPECT().DeleteCustomer(gomock.Any(), "1").Return(errors.New("some error"))
		res := testutil.DELETE(e, "/customer/1")
		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
}

func setupMock(t *testing.T) (*echo.Echo, func(), *mock_repository.MockCustomerRepository) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	close := func() {
		ctrl.Finish()
	}
	m := mock_repository.NewMockCustomerRepository(ctrl)
	h := NewDeleteCustomerHandler(*customerusecase.NewDeleteCustomerUseCase(m))
	h.RegisterRouter(e)
	return e, close, m
}

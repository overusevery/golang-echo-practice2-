package customerhandler

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	mock_repository "github.com/overusevery/golang-echo-practice2/src/repository/mock"
	"github.com/overusevery/golang-echo-practice2/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteCustomerHandler_DeleteCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		e := echo.New()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_repository.NewMockCustomerRepository(ctrl)
		h := NewDeleteCustomerHandler(m)
		h.RegisterRouter(e)

		m.EXPECT().DeleteCustomer(gomock.Any(), "1")
		res := testutil.GET(e, "/customer/1")
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
}

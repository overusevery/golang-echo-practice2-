package customerhandler

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCustomerHandler_DeleteCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		e := echo.New()
		h := &DeleteCustomerHandler{}
		h.RegisterRouter(e)
		res := testutil.GET(e, "/customer/1")
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
}

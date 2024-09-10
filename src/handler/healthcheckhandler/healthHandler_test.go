package healthcheckhandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		e := echo.New()
		h := &HealthHandler{}
		h.RegisterRouter(e)

		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		res := httptest.NewRecorder()
		e.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})
}

package healthcheckhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) RegisterRouter(e *echo.Echo) {
	e.GET("/health", h.HealthCheck)
}

func (h *HealthHandler) HealthCheck(c echo.Context) error {
	resMsg := HealthCheckResponse{
		Status: "OK",
	}

	return c.JSON(http.StatusOK, resMsg)
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

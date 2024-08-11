package customerhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DeleteCustomerHandler struct {
}

func NewDeleteCustomerHandler() *DeleteCustomerHandler {
	return &DeleteCustomerHandler{}
}

func (h *DeleteCustomerHandler) RegisterRouter(e *echo.Echo) {
	e.GET("/customer/:id", h.DeleteCustomer)
}

func (h *DeleteCustomerHandler) DeleteCustomer(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}

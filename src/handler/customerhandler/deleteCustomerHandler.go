package customerhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

type DeleteCustomerHandler struct {
	repository repository.CustomerRepository
}

func NewDeleteCustomerHandler(repository repository.CustomerRepository) *DeleteCustomerHandler {
	return &DeleteCustomerHandler{repository: repository}
}

func (h *DeleteCustomerHandler) RegisterRouter(e *echo.Echo) {
	e.GET("/customer/:id", h.DeleteCustomer)
}

func (h *DeleteCustomerHandler) DeleteCustomer(c echo.Context) error {
	id := c.Param("id")
	h.repository.DeleteCustomer(c.Request().Context(), id)
	return c.JSON(http.StatusOK, id)
}

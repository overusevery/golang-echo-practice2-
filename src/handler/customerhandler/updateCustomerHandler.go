package customerhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
)

type UpdateCustomerHandler struct {
	UpdateCustomerUseCase *customerusecase.UpdateCustomerUseCase
}

func NewUpdateCustomerHandler(UpdateCustomerUseCase *customerusecase.UpdateCustomerUseCase) *UpdateCustomerHandler {
	return &UpdateCustomerHandler{UpdateCustomerUseCase}
}

func (h *UpdateCustomerHandler) RegisterRouter(e *echo.Echo) {
	e.PUT("/customer/:id", h.UpdateCustomer)
}

func (h *UpdateCustomerHandler) UpdateCustomer(c echo.Context) error {
	id := c.Param("id")
	err := h.UpdateCustomerUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "ng")
	}

	return c.JSON(http.StatusOK, "ok")
}

package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
)

type CreateCustomerUseCase interface {
	Execute(ctx context.Context, customer entity.Customer) (*entity.Customer, error)
}
type CreateCustomerHandler struct {
	CreateCustomerUseCase CreateCustomerUseCase
}

func NewCreateCustomerHandler(createCustomerUseCase CreateCustomerUseCase) *CreateCustomerHandler {
	return &CreateCustomerHandler{createCustomerUseCase}
}

func (h *CreateCustomerHandler) RegisterRouter(e *echo.Echo) {
	e.POST("/customer", h.CreateCustomer)
}

func (h *CreateCustomerHandler) CreateCustomer(c echo.Context) error {
	customer := CreateCustomerRequest{}
	err := c.Bind(&customer)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	createdCustomer, err := h.CreateCustomerUseCase.Execute(c.Request().Context(), customer.ConvertFrom())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, convertToCreateCustomerResponse(*createdCustomer))
}

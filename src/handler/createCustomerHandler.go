package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase"
)

type CreateCustomerHandler struct {
	CreateCustomerUseCase usecase.CreateCustomerUseCase
}

func NewCreateCustomerHandler(createCustomerUseCase usecase.CreateCustomerUseCase) *CreateCustomerHandler {
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
	err = h.CreateCustomerUseCase.Execute(customer.ConvertFrom())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, CreateCustomerResponse{
		Status: "ok",
	})
}

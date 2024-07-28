package customerhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
)

type CreateCustomerHandler struct {
	CreateCustomerUseCase *customerusecase.CreateCustomerUseCase
}

func NewCreateCustomerHandler(createCustomerUseCase *customerusecase.CreateCustomerUseCase) *CreateCustomerHandler {
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

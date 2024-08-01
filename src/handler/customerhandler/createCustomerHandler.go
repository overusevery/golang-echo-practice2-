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
	req := CreateCustomerRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	customer, errList := req.ConvertFrom()
	if errList != nil {
		return c.JSON(http.StatusBadRequest, convertToCreateCustomerErrorResponse(errList))
	}
	createdCustomer, errList := h.CreateCustomerUseCase.Execute(c.Request().Context(), *customer)
	if errList != nil {
		return c.String(http.StatusInternalServerError, "bad request")
	}

	return c.JSON(http.StatusOK, convertToCreateCustomerResponse(*createdCustomer))
}

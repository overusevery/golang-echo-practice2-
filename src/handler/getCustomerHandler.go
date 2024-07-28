package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
)

type CustomerHandler struct {
	GetCustomerUseCase *customerusecase.GetCustomerUseCase
}

func NewGetCustomrHandler(getCustomerUseCase *customerusecase.GetCustomerUseCase) *CustomerHandler {
	return &CustomerHandler{getCustomerUseCase}
}

func (h *CustomerHandler) RegisterRouter(e *echo.Echo) {
	e.GET("/customer/:id", h.GetCustomer)
}

func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	customer, err := h.GetCustomerUseCase.Execute(c.Request().Context(), id)
	if err == repository.ErrCustomerNotFound {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Customer (id = %v) is not found", id))
	}
	if err != nil {
		return err
	}

	res := convertFrom(*customer)
	return c.JSON(http.StatusOK, res)
}

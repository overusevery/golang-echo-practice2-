package customerhandler

import (
	"errors"
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
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	customer, err := h.GetCustomerUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrCustomerNotFound):
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Customer (id = %v) is not found", id))
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	res := convertFrom(*customer)
	return c.JSON(http.StatusOK, res)
}

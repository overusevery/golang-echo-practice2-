package customerhandler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	"golang.org/x/exp/slog"
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
		return c.JSON(http.StatusBadRequest, convertToGetCustomerMultiResponse(err))
	}

	customer, err := h.GetCustomerUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrCustomerNotFound):
			return c.JSON(http.StatusNotFound, GetCustomerErrorResponse{
				Message: fmt.Sprintf("Customer (id = %v) is not found", id),
			})
		default:
			slog.Error("GetCustomer get unexpected error", "detail", err)
			return c.JSON(http.StatusInternalServerError, CreateCustomerErrorResponse{
				Message: "failed to get customer",
			})
		}
	}

	res := convertFrom(*customer)
	return c.JSON(http.StatusOK, res)
}

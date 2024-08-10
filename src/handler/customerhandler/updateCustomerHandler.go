package customerhandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
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
	req := CreateCustomerRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "ng")
	}
	customerRes, err := h.UpdateCustomerUseCase.Execute(
		c.Request().Context(),
		id,
		customerusecase.UpdateCustomerUseCaseInput{
			Name:          req.Name,
			Address:       req.Address,
			ZIP:           req.Zip,
			Phone:         req.Phone,
			MarketSegment: req.Mktsegment,
			Nation:        req.Nation,
			Birthdate:     req.Birthdate,
		},
	)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrCustomerNotFound):
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusInternalServerError, "ng")
	}

	return c.JSON(http.StatusOK, convertToUpdateCustomerResponse(*customerRes))
}

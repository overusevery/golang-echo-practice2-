package customerhandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	accesscontrol "github.com/overusevery/golang-echo-practice2/src/domain/usecase/accessControl"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	"golang.org/x/exp/slog"
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
	req := UpdateCustomerRequest{}
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
			Version:       req.Version,
		},
	)
	if err != nil {
		switch {
		case errors.Is(err, accesscontrol.ErrNotEnoughScope):
			return c.JSON(http.StatusUnauthorized, CreateCustomerErrorResponse{
				Message: "access token lacks needed scope",
			})
		case errors.Is(err, repository.ErrCustomerNotFound):
			return c.JSON(http.StatusNotFound, err)
		case errors.Is(err, repository.ErrConflict):
			return c.JSON(http.StatusConflict, err)
		default:
			slog.Error("UpdateCustomer Internal Server Error", "error", err, "request", req)
			return c.JSON(http.StatusInternalServerError, "ng")
		}
	}

	return c.JSON(http.StatusOK, convertToUpdateCustomerResponse(*customerRes))
}

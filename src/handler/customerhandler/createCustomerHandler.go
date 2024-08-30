package customerhandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	accesscontrol "github.com/overusevery/golang-echo-practice2/src/domain/usecase/accessControl"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	"golang.org/x/exp/slog"
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
		return c.JSON(http.StatusBadRequest, convertToCreateCustomerErrorResponse(err))
	}
	createdCustomer, err := h.CreateCustomerUseCase.Execute(
		c.Request().Context(),
		customerusecase.CreateCustomerUseCaseInput{
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
		case errors.Is(err, accesscontrol.ErrNotEnoughScope):
			return c.JSON(http.StatusUnauthorized, CreateCustomerErrorResponse{
				Message: "access token lacks needed scope",
			})
		case errors.Is(err, customerusecase.ErrInvalidInputCreateCustomerUseCase):
			return c.JSON(http.StatusBadRequest, convertToCreateCustomerErrorResponse(err))
		default:
			slog.Error("CreateCustomer get unexpected error", "detail", err)
			return c.JSON(http.StatusInternalServerError, CreateCustomerErrorResponse{
				Message: "failed to create customer",
			})
		}
	}

	return c.JSON(http.StatusOK, convertToCreateCustomerResponse(*createdCustomer))
}

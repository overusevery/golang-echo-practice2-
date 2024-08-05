package customerhandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
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
		errorList := util.NewValidationErrorList(err)
		return c.JSON(http.StatusBadRequest, convertToCreateCustomerErrorResponse(errorList))
	}

	createdCustomer, err := h.CreateCustomerUseCase.Execute(c.Request().Context(), customerusecase.CreateCustomerUseCaseInput{
		Name:          req.Name,
		Address:       req.Address,
		ZIP:           req.Zip,
		Phone:         req.Phone,
		MarketSegment: req.Mktsegment,
		Nation:        req.Nation,
		Birthdate:     req.Birthdate,
	})
	if err != nil {
		switch {
		case errors.Is(err, customerusecase.ErrInvalidInputCreateCustomerUseCase):
			return c.JSON(http.StatusBadRequest, convertToCreateCustomerErrorResponse(err))
		default:
			return c.String(http.StatusInternalServerError, "bad request")
		}
	}

	return c.JSON(http.StatusOK, convertToCreateCustomerResponse(*createdCustomer))
}

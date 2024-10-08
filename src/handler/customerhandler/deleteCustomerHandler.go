package customerhandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	accesscontrol "github.com/overusevery/golang-echo-practice2/src/domain/usecase/accessControl"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
)

type DeleteCustomerHandler struct {
	DeleteCustomerUseCase customerusecase.DeleteCustomerUseCase
}

func NewDeleteCustomerHandler(usecase customerusecase.DeleteCustomerUseCase) *DeleteCustomerHandler {
	return &DeleteCustomerHandler{DeleteCustomerUseCase: usecase}
}

func (h *DeleteCustomerHandler) RegisterRouter(e *echo.Echo) {
	e.DELETE("/customer/:id", h.DeleteCustomer)
}

func (h *DeleteCustomerHandler) DeleteCustomer(c echo.Context) error {
	id := c.Param("id")
	err := h.DeleteCustomerUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, accesscontrol.ErrNotEnoughScope):
			return c.JSON(http.StatusUnauthorized, CreateCustomerErrorResponse{
				Message: "access token lacks needed scope",
			})
		case errors.Is(err, repository.ErrCustomerNotFound):
			return c.JSON(http.StatusNotFound, id)
		}
		return c.JSON(http.StatusInternalServerError, id)
	}
	return c.JSON(http.StatusOK, id)
}

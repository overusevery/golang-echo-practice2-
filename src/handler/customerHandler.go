package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase"
)

type CustomerHandler struct {
}

func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	customer, err := usecase.GetCustomer(id)
	if err != nil {
		return err
	}

	res := convertFrom(customer)
	return c.JSON(http.StatusOK, res)
}

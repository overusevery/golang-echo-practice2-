package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	res := GetCustomerResponse{ID: id}
	return c.JSON(http.StatusOK, res)
}

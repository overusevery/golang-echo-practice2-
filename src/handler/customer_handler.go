package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCustomer(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

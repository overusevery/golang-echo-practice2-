package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase"
	"github.com/overusevery/golang-echo-practice2/src/handler"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	customerHandler := handler.CustomerHandler{GetCustomerUseCase: usecase.GetCustomerUseCase{}}
	customerHandler.RegisterRouter(e)
	//e.GET("/customer/:id", customerHandler.GetCustomer)

	// Start server
	s := http.Server{
		Addr:        ":1323",
		Handler:     e,
		ReadTimeout: 30 * time.Second,
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

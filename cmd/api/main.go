package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	"github.com/overusevery/golang-echo-practice2/src/handler"
	"github.com/overusevery/golang-echo-practice2/src/repository"
)

func main() {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres sslmode=disable", config.Host, config.User, config.Password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to DB!")
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	r := repository.NewRealCustomerRepository()
	customerHandler := handler.NewCustomrHandler(*customerusecase.NewGetCustomerUseCase(r))
	customerHandler.RegisterRouter(e)

	// Start server
	server := http.Server{
		Addr:        ":1323",
		Handler:     e,
		ReadTimeout: 30 * time.Second,
	}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type Config struct {
	Host     string
	User     string
	Password string
}

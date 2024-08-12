package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	handler "github.com/overusevery/golang-echo-practice2/src/handler/customerhandler"
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
	r := repository.NewRealCustomerRepository(db)
	getCustomerHandler := handler.NewGetCustomrHandler(customerusecase.NewGetCustomerUseCase(r))
	getCustomerHandler.RegisterRouter(e)

	createCustomerHandler := handler.NewCreateCustomerHandler(customerusecase.NewCreateCustomerUseCase(r))
	createCustomerHandler.RegisterRouter(e)

	updateCustomerHandler := handler.NewUpdateCustomerHandler(customerusecase.NewUpdateCustomerUseCase(r))
	updateCustomerHandler.RegisterRouter(e)

	deleteCustomerHandler := handler.NewDeleteCustomerHandler(*customerusecase.NewDeleteCustomerUseCase(r))
	deleteCustomerHandler.RegisterRouter(e)
	// Start server
	server := http.Server{
		Addr:        ":1323",
		Handler:     e,
		ReadTimeout: 30 * time.Second,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	fmt.Print("graceful shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	Host     string
	User     string
	Password string
}

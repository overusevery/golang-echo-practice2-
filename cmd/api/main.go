package main

import (
	"context"
	"database/sql"
	"errors"
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
	"github.com/overusevery/golang-echo-practice2/src/handler/customemiddleware"
	handler "github.com/overusevery/golang-echo-practice2/src/handler/customerhandler"
	healthHandler "github.com/overusevery/golang-echo-practice2/src/handler/healthcheckhandler"
	"github.com/overusevery/golang-echo-practice2/src/repository"
)

func main() {
	os.Exit(run())
}

func run() int {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Printf("failed to load config file:%v", err.Error())
		return 1
	}
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("failed to connect to postgresql:%v", err.Error())
		return 1
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Printf("failed to ping to postgresql:%v", pingErr.Error())
		return 1
	}
	log.Printf("Connected to DB!")
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//health check
	healthHandler.NewHealthHandler().RegisterRouter(e)

	//auth middleware
	e.Use(customemiddleware.ParseAuthorizationToken("/health"))
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

	serverclosed := make(chan error, 1)
	go func() {
		serverclosed <- server.ListenAndServe()
	}()

	select {
	case err := <-serverclosed:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server is closed by error:%v", err.Error())
			return 1
		}
	case <-ctx.Done():
		log.Printf("graceful shutdown...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			log.Print("failed to shutdown")
			return 1
		}
	}
	return 0
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
}

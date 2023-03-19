package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang-employee-crud-example/api/config"
	"golang-employee-crud-example/api/handler"
	"golang-employee-crud-example/repository"
	"log"
	"net/http"
	"os"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Connect to the database
	dbConfig := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), config.Database.Host, config.Database.Port, config.Database.Name)
	dbPool, err := pgxpool.Connect(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer dbPool.Close()

	// Create employee repository
	employeeRepo := repository.NewEmployeeRepository(dbPool)

	// Create employee handler
	employeeHandler := handler.NewEmployeeHandler(employeeRepo)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/employees", employeeHandler.GetEmployees)
	e.GET("/employees/:id", employeeHandler.GetEmployeeByID)
	e.POST("/employees", employeeHandler.CreateEmployee)
	e.PUT("/employees/:id", employeeHandler.UpdateEmployee)
	e.DELETE("/employees/:id", employeeHandler.DeleteEmployee)

	// Start server
	addr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	log.Printf("Server listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, e))
}

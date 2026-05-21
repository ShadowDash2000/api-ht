package main

import (
	"log"
	"net/http"
	"os"
	"rest-api-hitalent/internal/handler"
	departmentHandler "rest-api-hitalent/internal/handler/department"
	employeeHandler "rest-api-hitalent/internal/handler/employee"
	departmentRepo "rest-api-hitalent/internal/repository/department"
	employeeRepo "rest-api-hitalent/internal/repository/employee"
	departmentService "rest-api-hitalent/internal/service/department"
	employeeService "rest-api-hitalent/internal/service/employee"
	"rest-api-hitalent/pkg/recoverer"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	departmentRepository := departmentRepo.NewRepository(db)
	employeeRepository := employeeRepo.NewRepository(db)

	departmentService := departmentService.NewService(departmentRepository, employeeRepository)
	employeeService := employeeService.NewService(employeeRepository, departmentRepository)

	departmentHandler := departmentHandler.NewHandler(departmentService)
	employeeHandler := employeeHandler.NewHandler(employeeService)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux, departmentHandler, employeeHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      recoverer.Recoverer(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

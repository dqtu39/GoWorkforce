package main

import (
	"github.com/dqtu39/GoWorkforce/internal/delivery/http"
	"github.com/dqtu39/GoWorkforce/internal/repository"
	"github.com/dqtu39/GoWorkforce/internal/usecase"
	"github.com/dqtu39/GoWorkforce/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	dbConn, err := db.NewMySQLConn()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	employeeRepo := repository.NewEmployeeRepository(dbConn)
	employeeUsecase := usecase.NewEmployeeUseCase(employeeRepo)
	employeeHandler := http.NewEmployeeHandler(employeeUsecase)

	r := gin.Default()

	r.GET("/employees", employeeHandler.ListEmployee)
	r.GET("/employees/:id", employeeHandler.GetEmployee)
	r.POST("/employees", employeeHandler.CreateEmployee)
	r.PUT("/employees/:id", employeeHandler.UpdateEmployee)
	r.DELETE("/employees/:id", employeeHandler.DeleteEmployee)

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting server on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

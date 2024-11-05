package main

import (
	"github.com/dqtu39/GoWorkforce/internal/delivery/http"
	"github.com/dqtu39/GoWorkforce/internal/middleware"
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
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepo)
	employeeHandler := http.NewEmployeeHandler(employeeUseCase)

	userRepo := repository.NewUserRepository(dbConn)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := http.NewUserHandler(userUseCase)

	r := gin.Default()

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	employeeRoutes := r.Group("/employees")
	employeeRoutes.Use(middleware.JWTAuthMiddleware())
	{
		employeeRoutes.POST("", employeeHandler.CreateEmployee)
		employeeRoutes.GET("", employeeHandler.ListEmployee)
		employeeRoutes.GET("/:id", employeeHandler.GetEmployee)
		employeeRoutes.PUT("/:id", employeeHandler.UpdateEmployee)
		employeeRoutes.DELETE("/:id", employeeHandler.DeleteEmployee)
	}

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

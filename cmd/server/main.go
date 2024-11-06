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

func setupRoutes(r *gin.Engine, employeeHandler *http.EmployeeHandler, userHandler *http.UserHandler) {
	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// Auth routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		employees := api.Group("/employees")
		{
			employees.POST("", employeeHandler.CreateEmployee)
			employees.GET("", employeeHandler.ListEmployee)
			employees.GET("/:id", employeeHandler.GetEmployee)
			employees.PUT("/:id", employeeHandler.UpdateEmployee)
			employees.DELETE("/:id", employeeHandler.DeleteEmployee)
		}
	}
}

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

	setupRoutes(r, employeeHandler, userHandler)

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


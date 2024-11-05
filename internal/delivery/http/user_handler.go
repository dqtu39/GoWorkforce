package http

import (
	"github.com/dqtu39/GoWorkforce/internal/domain"
	"github.com/dqtu39/GoWorkforce/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase}
}

func (uh *UserHandler) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error in Register use case: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Starting registration process for user:", user.Username)
	err := uh.useCase.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("User registered successfully:", user.Username)
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (uh *UserHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := uh.useCase.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}

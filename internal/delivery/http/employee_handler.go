package http

import (
	"github.com/dqtu39/GoWorkforce/internal/domain"
	"github.com/dqtu39/GoWorkforce/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	usecase usecase.EmployeeUseCase
}

func NewEmployeeHandler(userUseCase usecase.EmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{usecase: userUseCase}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var employee domain.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.usecase.CreateEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	employee.ID = int64(id)
	c.JSON(http.StatusCreated, gin.H{"employee": employee})
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var employee domain.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := h.usecase.UpdateEmployee(id, employee)
	if err != nil {
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"employee": employee})
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	rowsAffected, err := h.usecase.DeleteEmployee(id)
	if err != nil {
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "employee deleted"})
}

func (h *EmployeeHandler) ListEmployee(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
		return
	}

	employees, err := h.usecase.GetAllEmployee(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"employees": employees})
}

func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	employee, err := h.usecase.GetEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"employee": employee})
}

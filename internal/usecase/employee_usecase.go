package usecase

import (
	"github.com/dqtu39/GoWorkforce/internal/domain"
	"github.com/dqtu39/GoWorkforce/internal/repository"
)

type EmployeeUseCase interface {
	GetAllEmployee(offset, limit int) ([]domain.Employee, error)
	GetEmployeeById(id int) (domain.Employee, error)
	CreateEmployee(emp domain.Employee) (int64, error)
	UpdateEmployee(id int, emp domain.Employee) (int64, error)
	DeleteEmployee(id int) (int64, error)
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo}
}

func (uc *employeeUseCase) GetAllEmployee(offset, limit int) ([]domain.Employee, error) {
	return uc.repo.GetAll(offset, limit)
}

func (uc *employeeUseCase) GetEmployeeById(id int) (domain.Employee, error) {
	return uc.repo.GetEmployee(id)
}

func (uc *employeeUseCase) CreateEmployee(emp domain.Employee) (int64, error) {
	return uc.repo.CreateEmployee(emp)
}

func (uc *employeeUseCase) UpdateEmployee(id int, emp domain.Employee) (int64, error) {
	return uc.repo.UpdateEmployee(id, emp)
}

func (uc *employeeUseCase) DeleteEmployee(id int) (int64, error) {
	return uc.repo.DeleteEmployee(id)
}

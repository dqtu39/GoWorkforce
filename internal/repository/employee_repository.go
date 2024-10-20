package repository

import (
	"database/sql"
	"errors"
	"github.com/dqtu39/GoWorkforce/internal/domain"
)

var (
	ErrNotFound = errors.New("employee not found")
)

type EmployeeRepository interface {
	GetAll(offset int, limit int) ([]domain.Employee, error)
	GetEmployee(id int) (domain.Employee, error)
	CreateEmployee(employee domain.Employee) (int64, error)
	UpdateEmployee(id int, employee domain.Employee) (int64, error)
	DeleteEmployee(id int) (int64, error)
}

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) GetAll(offset int, limit int) ([]domain.Employee, error) {
	rows, err := r.db.Query("SELECT * FROM employees LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	employees := make([]domain.Employee, 0)
	for rows.Next() {
		var employee domain.Employee
		if err := rows.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Age, &employee.Position); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (r *employeeRepository) GetEmployee(id int) (domain.Employee, error) {
	employee := domain.Employee{}
	err := r.db.QueryRow("SELECT * FROM employees WHERE ID = ?", id).Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Age, &employee.Position)
	if err == sql.ErrNoRows {
		return domain.Employee{}, ErrNotFound
	}
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

func (r *employeeRepository) CreateEmployee(employee domain.Employee) (int64, error) {
	res, err := r.db.Exec("INSERT INTO employees (first_name, last_name, age, position) VALUES (?,?,?,?)", employee.FirstName, employee.LastName, employee.Age, employee.Position)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *employeeRepository) UpdateEmployee(id int, employee domain.Employee) (int64, error) {
	// Check if the employee exists before trying to update
	_, err := r.GetEmployee(id)
	if err != nil {
		if err == ErrNotFound {
			return 0, ErrNotFound
		}
		return 0, err
	}

	// Update the employee record
	res, err := r.db.Exec("UPDATE employees SET first_name = ?, last_name = ?, age = ?, position = ? WHERE id = ?", employee.FirstName, employee.LastName, employee.Age, employee.Position, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *employeeRepository) DeleteEmployee(id int) (int64, error) {
	// Check if the employee exists before trying to delete
	_, err := r.GetEmployee(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return 0, ErrNotFound
		}
		return 0, err
	}

	// Delete the employee record
	res, err := r.db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

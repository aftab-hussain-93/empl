package types

import "context"

type EmployeeService interface {
	CreateEmployee(context.Context, *Employee) (*Employee, error)
	GetEmployeeByID(context.Context, uint) (*Employee, error)
	UpdateEmployee(context.Context, uint, *Employee) (*Employee, error)
	DeleteEmployee(context.Context, uint) error
}

type Employee struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

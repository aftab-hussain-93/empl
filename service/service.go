package service

import (
	"context"

	"github.com/aftab-hussain-93/empl/types"
)

type Repository interface {
	Create(context.Context, *types.Employee) (*types.Employee, error)
	Update(context.Context, uint, *types.Employee) (*types.Employee, error)
	GetByID(context.Context, uint) (*types.Employee, error)
	Get(context.Context, int64, int64) ([]*types.Employee, error)
}

type employeeSvc struct {
	repo Repository
}

// CreateEmployee implements EmployeeService.
func (e *employeeSvc) CreateEmployee(context.Context, *types.Employee) (*types.Employee, error) {
	panic("unimplemented")
}

// DeleteEmployee implements EmployeeService.
func (e *employeeSvc) DeleteEmployee(context.Context, uint) error {
	panic("unimplemented")
}

// GetEmployeeByID implements EmployeeService.
func (e *employeeSvc) GetEmployeeByID(context.Context, uint) (*types.Employee, error) {
	panic("unimplemented")
}

// UpdateEmployee implements EmployeeService.
func (e *employeeSvc) UpdateEmployee(context.Context, uint, *types.Employee) (*types.Employee, error) {
	panic("unimplemented")
}

var _ types.EmployeeService = (*employeeSvc)(nil)

func NewService(repo Repository) *employeeSvc {
	return &employeeSvc{repo}
}

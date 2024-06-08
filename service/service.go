package service

import (
	"context"
)

type employeeSvc struct {
	repo Repository
}

// CreateEmployee implements EmployeeService.
func (s *employeeSvc) CreateEmployee(ctx context.Context, emp *EmployeeCreateRequest) (*Employee, error) {
	// return s.repo.Create(ctx, emp)
	panic("unimplemented")
}

// DeleteEmployee implements EmployeeService.
func (e *employeeSvc) DeleteEmployee(context.Context, uint) error {
	panic("unimplemented")
}

// GetEmployeeByID implements EmployeeService.
func (e *employeeSvc) GetEmployeeByID(context.Context, uint) (*Employee, error) {
	panic("unimplemented")
}

// UpdateEmployee implements EmployeeService.
func (e *employeeSvc) UpdateEmployee(context.Context, uint, *Employee) (*Employee, error) {
	panic("unimplemented")
}

var _ EmployeeService = (*employeeSvc)(nil)

func NewService(repo Repository) *employeeSvc {
	return &employeeSvc{repo}
}

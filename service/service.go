package service

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type employeeSvc struct {
	repo Repository
}

// CreateEmployee implements EmployeeService.
func (s *employeeSvc) CreateEmployee(ctx context.Context, req *EmployeeCreateRequest) (*Employee, error) {
	emp := &Employee{
		Name:     req.Name,
		Position: req.Position,
		Salary:   req.Salary,
	}
	emp, err := s.repo.Create(ctx, emp)
	if err != nil {
		return nil, err
	}

	return emp, nil
}

// DeleteEmployee implements EmployeeService.
func (e *employeeSvc) DeleteEmployee(context.Context, uint) error {
	panic("unimplemented")
}

// GetEmployeeByID implements EmployeeService.
func (e *employeeSvc) GetEmployeeByID(context.Context, uint) (*Employee, error) {
	panic("unimplemented")
}

// GetEmployeeByID implements EmployeeService.
func (s *employeeSvc) ListEmployees(ctx context.Context, pageSize, pageNumber int64) ([]*Employee, int64, error) {
	offset := (pageNumber - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	limit := pageSize
	if limit > MAX_PAGE_SIZE || limit <= 0 {
		limit = MAX_PAGE_SIZE
	}

	page := Pagination{
		Offset: offset,
		Limit:  limit,
	}

	var count int64 = 0
	var resp []*Employee

	// getting total count and rows
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		cnt, err := s.repo.Count(egCtx)
		if err != nil {
			return err
		}
		count = cnt

		return nil
	})

	eg.Go(func() error {
		emps, err := s.repo.Get(ctx, page)
		if err != nil {
			return err
		}
		resp = emps
		return nil
	})

	if err := eg.Wait(); err != nil {
		return resp, count, err
	}

	return resp, count, nil
}

// UpdateEmployee implements EmployeeService.
func (e *employeeSvc) UpdateEmployee(context.Context, uint, *Employee) (*Employee, error) {
	panic("unimplemented")
}

var _ EmployeeService = (*employeeSvc)(nil)

func NewService(repo Repository) *employeeSvc {
	return &employeeSvc{repo}
}

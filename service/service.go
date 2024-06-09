package service

import (
	"context"
	"database/sql"
	"errors"

	fault "github.com/aftab-hussain-93/empl/err"
	"github.com/jackc/pgx/v5"
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
func (s *employeeSvc) DeleteEmployee(ctx context.Context, id int) error {
	// check if employee exists
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		// return 404 error is not found
		if errors.Is(err, sql.ErrNoRows) {
			return fault.New(fault.ErrNotFound, "Record not found", nil)
		}
		return fault.New(fault.ErrInternalServer, "Internal Server Error", err)
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return fault.New(fault.ErrInternalServer, "Internal Server Error", err)
	}

	return nil
}

// GetEmployeeByID implements EmployeeService.
func (s *employeeSvc) GetEmployeeByID(ctx context.Context, id int) (*Employee, error) {
	// check if employee exists
	emp, err := s.repo.GetByID(ctx, id)
	if err != nil {
		// return 404 error is not found
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fault.New(fault.ErrNotFound, "Record not found", nil)
		}
		return nil, fault.New(fault.ErrInternalServer, "Internal Server Error", err)
	}

	return emp, nil
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
		emps, err := s.repo.Get(egCtx, page)
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
func (s *employeeSvc) UpdateEmployee(ctx context.Context, id int, data *EmployeeUpdateRequest) (*Employee, error) {
	// check if employee exists
	emp, err := s.repo.GetByID(ctx, id)
	if err != nil {
		// return 404 error is not found
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fault.New(fault.ErrNotFound, "Record not found", nil)
		}
		return nil, fault.New(fault.ErrInternalServer, "Internal Server Error", err)
	}

	// setting the updated fields
	if data.Name != "" {
		emp.Name = data.Name
	}
	if data.Position != "" {
		emp.Position = data.Position
	}
	if data.Salary > 0 {
		emp.Salary = data.Salary
	}

	emp, err = s.repo.Update(ctx, id, emp)
	if err != nil {
		return nil, fault.New(fault.ErrInternalServer, "Internal Server Error", err)
	}

	return emp, nil
}

var _ EmployeeService = (*employeeSvc)(nil)

func NewService(repo Repository) *employeeSvc {
	return &employeeSvc{repo}
}

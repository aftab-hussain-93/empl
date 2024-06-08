package repository

import (
	"context"
	"fmt"

	"github.com/aftab-hussain-93/empl/service"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

// Create implements Repository.
func (r *repository) Create(ctx context.Context, emp *service.Employee) (*service.Employee, error) {
	query := `INSERT INTO employees (name, position, salary) VALUES (@name, @position, @salary) RETURNING id, name, position, salary;`
	args := pgx.NamedArgs{
		"name":     emp.Name,
		"position": emp.Position,
		"salary":   emp.Salary,
	}
	res := service.Employee{}
	err := r.db.QueryRow(ctx, query, args).Scan(&res.ID, &res.Name, &res.Position, &res.Salary)
	if err != nil {
		return nil, fmt.Errorf("unable to insert row: %w", err)
	}

	return &res, nil
}

// Get implements Repository.
func (r *repository) Get(ctx context.Context, pg service.Pagination) ([]*service.Employee, error) {
	query := `SELECT id, name, position, salary FROM employees OFFSET @offset LIMIT @limit;`
	args := pgx.NamedArgs{
		"offset": pg.Offset,
		"limit":  pg.Limit,
	}
	rows, err := r.db.Query(ctx, query, args)
	if err != nil {
		return []*service.Employee{}, fmt.Errorf("unable to get employees %w", err)
	}
	defer rows.Close()

	resp := []*service.Employee{}

	for rows.Next() {
		row := &service.Employee{}
		if err := rows.Scan(&row.ID, &row.Name, &row.Position, &row.Salary); err != nil {
			return []*service.Employee{}, fmt.Errorf("unable to scan row %w", err)
		}
		resp = append(resp, row)
	}

	return resp, nil
}

func (r *repository) Count(ctx context.Context) (int64, error) {
	query := `SELECT count(id) FROM employees;`
	var cnt int64
	err := r.db.QueryRow(ctx, query).Scan(&cnt)
	if err != nil {
		return 0, fmt.Errorf("unable to get rows count: %w", err)
	}

	return cnt, nil
}

// GetByID implements Repository.
func (r *repository) GetByID(context.Context, uint) (*service.Employee, error) {
	panic("unimplemented")
}

// Update implements Repository.
func (r *repository) Update(context.Context, uint, *service.Employee) (*service.Employee, error) {
	panic("unimplemented")
}

var _ service.Repository = (*repository)(nil)

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{db}
}

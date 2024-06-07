package repository

import (
	"context"

	"github.com/aftab-hussain-93/empl/service"
	"github.com/aftab-hussain-93/empl/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

// Create implements Repository.
func (r *repository) Create(context.Context, *types.Employee) (*types.Employee, error) {
	panic("unimplemented")
}

// Get implements Repository.
func (r *repository) Get(context.Context, int64, int64) ([]*types.Employee, error) {
	panic("unimplemented")
}

// GetByID implements Repository.
func (r *repository) GetByID(context.Context, uint) (*types.Employee, error) {
	panic("unimplemented")
}

// Update implements Repository.
func (r *repository) Update(context.Context, uint, *types.Employee) (*types.Employee, error) {
	panic("unimplemented")
}

var _ service.Repository = (*repository)(nil)

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{db}
}

package service

import "context"

type Repository interface {
	Create(context.Context, *Employee) (*Employee, error)
	Update(context.Context, uint, *Employee) (*Employee, error)
	GetByID(context.Context, uint) (*Employee, error)
	Get(context.Context, Pagination) ([]*Employee, error)
	Count(ctx context.Context) (int64, error)
}

type EmployeeService interface {
	CreateEmployee(context.Context, *EmployeeCreateRequest) (*Employee, error)
	GetEmployeeByID(context.Context, uint) (*Employee, error)
	UpdateEmployee(context.Context, uint, *Employee) (*Employee, error)
	DeleteEmployee(context.Context, uint) error
	ListEmployees(ctx context.Context, pageSize int64, pageNumber int64) ([]*Employee, int64, error)
}

type Pagination struct {
	Offset int64
	Limit  int64
}

type Employee struct {
	ID       uint             `json:"id"`
	Name     string           `json:"name"`
	Position EmployeePosition `json:"position"`
	Salary   float64          `json:"salary"`
}

type EmployeeCreateRequest struct {
	Name     string           `json:"name" validate:"required,max=50"`
	Position EmployeePosition `json:"position" validate:"required,oneof=manager trainee"`
	Salary   float64          `json:"salary" validate:"gt=0"`
}

type EmployeeUpdateRequest struct {
	Name     string           `json:"name" validate:"omitempty,max=50"`
	Position EmployeePosition `json:"position" validate:"omitempty,oneof=manager trainee"`
	Salary   float64          `json:"salary" validate:"omitempty,gt=0"`
}

type EmployeePosition string

const (
	EmployeePositionManager EmployeePosition = "manager"
	EmployeePositionTrainee EmployeePosition = "trainee"
)

const MAX_PAGE_SIZE = 100

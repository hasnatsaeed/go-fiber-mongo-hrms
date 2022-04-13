package repositories

import (
	"context"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/models"
)

type Repository interface {
	SaveEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error)
	FetchEmployeeById(ctx context.Context, id string) (*models.Employee, error)
	FetchAllEmployees(ctx context.Context) ([]models.Employee, error)
	UpdateEmployee(ctx context.Context, id string, update *models.Employee) (*models.Employee, error)
	DeleteEmployee(ctx context.Context, id string) (*models.Employee, error)
}

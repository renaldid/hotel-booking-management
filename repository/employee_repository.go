package repository

import (
	"context"
	"database/sql"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type EmployeeRepository interface {
	Create(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee
	Update(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee
	Delete(ctx context.Context, tx *sql.Tx, employee domain.Employee)
	FindById(ctx context.Context, tx *sql.Tx, employeeId int) (domain.Employee, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Employee
}

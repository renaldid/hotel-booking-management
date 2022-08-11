package repository

import (
	"context"
	"database/sql"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type RoleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.Role) domain.Role
	Update(ctx context.Context, tx *sql.Tx, user domain.Role) domain.Role
	Delete(ctx context.Context, tx *sql.Tx, user domain.Role)
	FindById(ctx context.Context, tx *sql.Tx, roleId int) (domain.Role, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Role
}

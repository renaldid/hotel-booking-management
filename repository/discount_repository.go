package repository

import (
	"context"
	"database/sql"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type DiscountRepository interface {
	Create(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount
	Update(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount
	Delete(ctx context.Context, tx *sql.Tx, discount domain.Discount)
	FindById(ctx context.Context, tx *sql.Tx, discountId int) (domain.Discount, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Discount
}

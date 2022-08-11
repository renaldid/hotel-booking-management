package repository

import (
	"context"
	"database/sql"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type BookingRepository interface {
	Create(ctx context.Context, tx *sql.Tx, booking domain.Booking) domain.Booking
	Update(ctx context.Context, tx *sql.Tx, booking domain.Booking) domain.Booking
	Delete(ctx context.Context, tx *sql.Tx, booking domain.Booking)
	FindById(ctx context.Context, tx *sql.Tx, bookingId int) (domain.Booking, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Booking
}

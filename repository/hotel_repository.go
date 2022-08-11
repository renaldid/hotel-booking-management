package repository

import (
	"context"
	"database/sql"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type HotelRepository interface {
	Create(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel
	Update(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel
	Delete(ctx context.Context, tx *sql.Tx, hotel domain.Hotel)
	FindById(ctx context.Context, tx *sql.Tx, hotelId int) (domain.Hotel, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Hotel
}

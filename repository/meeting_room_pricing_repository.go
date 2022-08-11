package repository

import (
	"context"
	"database/sql"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type MeetingRoomPricingRepository interface {
	Create(ctx context.Context, tx *sql.Tx, meetingRoomPricing domain.MeetingRoomPricing) domain.MeetingRoomPricing
	Update(ctx context.Context, tx *sql.Tx, meetingRoomPricing domain.MeetingRoomPricing) domain.MeetingRoomPricing
	Delete(ctx context.Context, tx *sql.Tx, meetingRoomPricing domain.MeetingRoomPricing)
	FindById(ctx context.Context, tx *sql.Tx, meetingRoomPricingId int) (domain.MeetingRoomPricing, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MeetingRoomPricing
}

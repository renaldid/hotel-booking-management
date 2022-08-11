package repository

import (
	"context"
	"database/sql"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type MeetingRoomRepository interface {
	Create(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom) domain.MeetingRoom
	Update(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom) domain.MeetingRoom
	Delete(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom)
	FindById(ctx context.Context, tx *sql.Tx, meetingRoomId int) (domain.MeetingRoom, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MeetingRoom
}

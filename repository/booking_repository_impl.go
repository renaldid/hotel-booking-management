package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type BookingRepositoryImpl struct {
}

func NewBookingRepository() BookingRepository {
	return &BookingRepositoryImpl{}
}

func (repository BookingRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, booking domain.Booking) domain.Booking {
	SQL := "insert into bookings(meeting_room_id, invoice_id, status) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, booking.MeetingRoomId, booking.InvoiceId, booking.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	booking.Id = int(id)
	return booking
}

func (repository BookingRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, booking domain.Booking) domain.Booking {
	SQL := "update bookings set meeting_room_id = ?, invoice_id = ?, status = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, booking.MeetingRoomId, booking.InvoiceId, booking.Status, booking.Id)
	helper.PanicIfError(err)

	return booking
}

func (repository BookingRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, booking domain.Booking) {
	SQL := "delete from bookings where id = ?"
	_, err := tx.ExecContext(ctx, SQL, booking.Id)
	helper.PanicIfError(err)
}

func (repository BookingRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookingId int) (domain.Booking, error) {
	SQL := "select b.id, b.meeting_room_id, b.invoice_id, b.status, m.name as meeting_room_name, i.number as invoice_number, i.pic as pic_name from ((bookings b inner join meeting_rooms m on b.meeting_room_id=m.id) inner join invoices i on b.invoice_id=i.id) where b.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, bookingId)
	helper.PanicIfError(err)
	defer rows.Close()

	booking := domain.Booking{}
	if rows.Next() {
		err := rows.Scan(&booking.Id, &booking.MeetingRoomId, &booking.InvoiceId, &booking.Status, &booking.MeetingRoomName, &booking.InvoiceNumber, &booking.PicName)
		helper.PanicIfError(err)
		return booking, nil
	} else {
		return booking, errors.New("booking is not found")
	}
}

func (repository BookingRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Booking {
	SQL := "select b.id, b.meeting_room_id, b.invoice_id, b.status, m.name as meeting_room_name, i.number as invoice_number, i.pic as pic_name from ((bookings b inner join meeting_rooms m on b.meeting_room_id=m.id) inner join invoices i on b.invoice_id=i.id) "
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var bookings []domain.Booking
	for rows.Next() {
		booking := domain.Booking{}
		err := rows.Scan(&booking.Id, &booking.MeetingRoomId, &booking.InvoiceId, &booking.Status, &booking.MeetingRoomName, &booking.InvoiceNumber, &booking.PicName)
		helper.PanicIfError(err)
		bookings = append(bookings, booking)
	}
	return bookings
}

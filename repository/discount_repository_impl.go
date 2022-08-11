package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type DiscountRepositoryImpl struct {
}

func NewDiscountRepository() DiscountRepository {
	return &DiscountRepositoryImpl{}
}

func (repository DiscountRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount {
	SQL := "insert into discounts(employee_id, hotel_id, meeting_room_id, rate, status, request_date) values (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, discount.EmployeeId, discount.HotelId, discount.MeetingRoomId, discount.Rate, discount.Status, discount.RequestDate)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	discount.Id = int(id)
	return discount
}

func (repository DiscountRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount {
	SQL := "update discounts set employee_id =?, hotel_id=?,meeting_room_id=?,rate = ?, status = ?, request_date = ?   where id = ?"
	_, err := tx.ExecContext(ctx, SQL, discount.EmployeeId, discount.HotelId, discount.MeetingRoomId, discount.Rate, discount.Status, discount.RequestDate, discount.Id)
	helper.PanicIfError(err)

	return discount
}

func (repository DiscountRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, discount domain.Discount) {
	SQL := "delete from discounts where id = ?"
	_, err := tx.ExecContext(ctx, SQL, discount.Id)
	helper.PanicIfError(err)
}

func (repository DiscountRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, discountId int) (domain.Discount, error) {
	SQL := "SELECT d.id, d.employee_id, d.hotel_id, d.meeting_room_id, d.rate, d.status, d.request_date, e.name as employee_name, h.name as hotel_name, e.name as meeting_room_name " +
		"from discounts d INNER JOIN employees e on d.employee_id=e.id INNER JOIN hotels h on d.hotel_id=h.id INNER JOIN meeting_rooms m on d.meeting_room_id=m.id where d.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, discountId)
	helper.PanicIfError(err)
	defer rows.Close()

	discount := domain.Discount{}
	if rows.Next() {
		err := rows.Scan(&discount.Id, &discount.EmployeeId, &discount.HotelId, &discount.MeetingRoomId, &discount.Rate, &discount.Status, &discount.RequestDate, &discount.EmployeeName, &discount.HotelName, &discount.RoomName)
		helper.PanicIfError(err)
		return discount, nil
	} else {
		return discount, errors.New("id not found")
	}
}

func (repository DiscountRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Discount {
	SQL := "SELECT d.id, d.employee_id, d.hotel_id, d.meeting_room_id, d.rate, d.status, d.request_date, e.name as employee_name, h.name as hotel_name, e.name as meeting_room_name " +
		"from discounts d INNER JOIN employees e on d.employee_id=e.id INNER JOIN hotels h on d.hotel_id=h.id INNER JOIN meeting_rooms m on d.meeting_room_id=m.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var discounts []domain.Discount
	for rows.Next() {
		discount := domain.Discount{}
		err := rows.Scan(&discount.Id, &discount.EmployeeId, &discount.HotelId, &discount.MeetingRoomId, &discount.Rate, &discount.Status, &discount.RequestDate, &discount.EmployeeName, &discount.HotelName, &discount.RoomName)
		helper.PanicIfError(err)
		discounts = append(discounts, discount)
	}
	return discounts
}

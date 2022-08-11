package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type MeetingRoomPricingRepositoryImpl struct {
}

func NewMeetingRoomPricingRepository() MeetingRoomPricingRepository {
	return &MeetingRoomPricingRepositoryImpl{}
}

func (repository MeetingRoomPricingRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, meetingRoomPricing domain.MeetingRoomPricing) domain.MeetingRoomPricing {
	SQL := "insert into meeting_room_pricings(meeting_room_id, price, price_type) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, meetingRoomPricing.MeetingRoomId, meetingRoomPricing.Price, meetingRoomPricing.PriceType)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	meetingRoomPricing.Id = int(id)
	return meetingRoomPricing
}

func (repository MeetingRoomPricingRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, meetingRoomPricing domain.MeetingRoomPricing) domain.MeetingRoomPricing {
	SQL := "update meeting_room_pricings set meeting_room_id = ?, price = ?, price_type = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, meetingRoomPricing.MeetingRoomId, meetingRoomPricing.Price, meetingRoomPricing.PriceType, meetingRoomPricing.Id)
	helper.PanicIfError(err)

	return meetingRoomPricing
}

func (repository MeetingRoomPricingRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, meetingRoomPricing domain.MeetingRoomPricing) {
	SQL := "delete from meeting_room_pricings where id = ?"
	_, err := tx.ExecContext(ctx, SQL, meetingRoomPricing.Id)
	helper.PanicIfError(err)
}

func (repository MeetingRoomPricingRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, meetingRoomPricingId int) (domain.MeetingRoomPricing, error) {
	SQL := "select p.id, p.meeting_room_id, p.price, p.price_type, m.name , m.capacity  from meeting_room_pricings p INNER JOIN meeting_rooms m on p.meeting_room_id= m.id  where p.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, meetingRoomPricingId)
	helper.PanicIfError(err)
	defer rows.Close()

	meetingRoomPricing := domain.MeetingRoomPricing{}
	if rows.Next() {
		err := rows.Scan(&meetingRoomPricing.Id, &meetingRoomPricing.MeetingRoomId, &meetingRoomPricing.Price, &meetingRoomPricing.PriceType, &meetingRoomPricing.Name, &meetingRoomPricing.Capacity)
		helper.PanicIfError(err)
		return meetingRoomPricing, nil
	} else {
		return meetingRoomPricing, errors.New("Id is not found")
	}
}

func (repository MeetingRoomPricingRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MeetingRoomPricing {
	SQL := "select p.id, p.meeting_room_id, p.price, p.price_type, m.name , m.capacity  from meeting_room_pricings p INNER JOIN meeting_rooms m on p.meeting_room_id= m.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var meetingRoomPricings []domain.MeetingRoomPricing
	for rows.Next() {
		meetingRoomPricing := domain.MeetingRoomPricing{}
		err := rows.Scan(&meetingRoomPricing.Id, &meetingRoomPricing.MeetingRoomId, &meetingRoomPricing.Price, &meetingRoomPricing.PriceType, &meetingRoomPricing.Name, &meetingRoomPricing.Capacity)
		helper.PanicIfError(err)
		meetingRoomPricings = append(meetingRoomPricings, meetingRoomPricing)
	}
	return meetingRoomPricings
}

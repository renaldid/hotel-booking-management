package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type MeetingRoomRepositoryImpl struct {
}

func NewMeetingRoomRepository() MeetingRoomRepository {
	return &MeetingRoomRepositoryImpl{}
}

func (repository MeetingRoomRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom) domain.MeetingRoom {
	SQL := "insert into meeting_rooms(floor_id, name, capacity, facility_id) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, meetingRoom.FloorId, meetingRoom.Name, meetingRoom.Capacity, meetingRoom.FacilityId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	meetingRoom.Id = int(id)
	return meetingRoom
}

func (repository MeetingRoomRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom) domain.MeetingRoom {
	SQL := "update meeting_rooms set floor_id = ?, name = ?, capacity = ?, facility_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, meetingRoom.FloorId, meetingRoom.Name, meetingRoom.Capacity, meetingRoom.FacilityId, meetingRoom.Id)
	helper.PanicIfError(err)

	return meetingRoom
}

func (repository MeetingRoomRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom) {
	SQL := "delete from meeting_rooms where id = ?"
	_, err := tx.ExecContext(ctx, SQL, meetingRoom.Id)
	helper.PanicIfError(err)
}

func (repository MeetingRoomRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, meetingRoomId int) (domain.MeetingRoom, error) {
	SQL := "select m.id, m.floor_id, m.name, m.capacity, m.facility_id, a.name as floor_name, a.description as floor_description, f.name as faccility_name, f.description as facility_description from meeting_rooms m INNER JOIN floors a on m.floor_id= a.id INNER JOIN facilities f on m.facility_id= f.id where m.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, meetingRoomId)
	helper.PanicIfError(err)
	defer rows.Close()

	meetingRoom := domain.MeetingRoom{}
	if rows.Next() {
		err := rows.Scan(&meetingRoom.Id, &meetingRoom.FloorId, &meetingRoom.Name, &meetingRoom.Capacity, &meetingRoom.FacilityId, &meetingRoom.FloorName, &meetingRoom.FloorDescription, &meetingRoom.FacilityName, &meetingRoom.FacilityDescription)
		helper.PanicIfError(err)
		return meetingRoom, nil
	} else {
		return meetingRoom, errors.New("Id is not found")
	}
}

func (repository MeetingRoomRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MeetingRoom {
	SQL := "select m.id, m.floor_id, m.name, m.capacity, m.facility_id, a.name as floor_name, a.description as floor_description, f.name as facility_name, f.description as facility_description from meeting_rooms m INNER JOIN floors a on m.floor_id= a.id INNER JOIN facilities f on m.facility_id= f.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var meetingRooms []domain.MeetingRoom
	for rows.Next() {
		meetingRoom := domain.MeetingRoom{}
		err := rows.Scan(&meetingRoom.Id, &meetingRoom.FloorId, &meetingRoom.Name, &meetingRoom.Capacity, &meetingRoom.FacilityId, &meetingRoom.FloorName, &meetingRoom.FloorDescription, &meetingRoom.FacilityName, &meetingRoom.FacilityDescription)
		helper.PanicIfError(err)
		meetingRooms = append(meetingRooms, meetingRoom)
	}
	return meetingRooms
}

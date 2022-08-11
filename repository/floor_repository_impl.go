package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type FloorRepositoryImpl struct {
}

func NewFloorRepository() FloorRepository {
	return &FloorRepositoryImpl{}
}

func (repository FloorRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor {
	SQL := "insert into floors(hotel_id, name, description) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, floor.HotelId, floor.Name, floor.Description)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	floor.Id = int(id)
	return floor
}

func (repository FloorRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor {
	SQL := "update floors set hotel_id = ?, name = ?, description = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, floor.HotelId, floor.Name, floor.Description, floor.Id)
	helper.PanicIfError(err)

	return floor
}

func (repository FloorRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, floor domain.Floor) {
	SQL := "delete from floors where id = ?"
	_, err := tx.ExecContext(ctx, SQL, floor.Id)
	helper.PanicIfError(err)
}

func (repository FloorRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, floorId int) (domain.Floor, error) {
	SQL := "select a.id, a.hotel_id, a.name, a.description, h.name as hotel_name from floors a inner join hotels h on a.hotel_id=h.id where a.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, floorId)
	helper.PanicIfError(err)
	defer rows.Close()

	floor := domain.Floor{}
	if rows.Next() {
		err := rows.Scan(&floor.Id, &floor.HotelId, &floor.Name, &floor.Description, &floor.HotelName)
		helper.PanicIfError(err)
		return floor, nil
	} else {
		return floor, errors.New("floor is not found")
	}
}

func (repository FloorRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Floor {
	SQL := "select a.id, a.hotel_id, a.name, a.description, h.name as hotel_name from floors a inner join hotels h on a.hotel_id=h.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var floors []domain.Floor
	for rows.Next() {
		floor := domain.Floor{}
		err := rows.Scan(&floor.Id, &floor.HotelId, &floor.Name, &floor.Description, &floor.HotelName)
		helper.PanicIfError(err)
		floors = append(floors, floor)
	}
	return floors
}

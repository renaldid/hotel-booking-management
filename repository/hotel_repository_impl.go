package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type HotelRepositoryImpl struct {
}

func NewHotelRepository() HotelRepository {
	return &HotelRepositoryImpl{}
}

func (repository HotelRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel {
	SQL := "insert into hotels(name, address, city, zip_code, rate) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, hotel.Name, hotel.Address, hotel.City, hotel.ZipCode, hotel.Rate)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	hotel.Id = int(id)
	return hotel
}

func (repository HotelRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel {
	SQL := "update hotels set name = ?, address = ?, city= ?, zip_code=?, rate=? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, hotel.Name, hotel.Address, hotel.City, hotel.ZipCode, hotel.Rate, hotel.Id)
	helper.PanicIfError(err)

	return hotel
}

func (repository HotelRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) {
	SQL := "delete from hotels where id = ?"
	_, err := tx.ExecContext(ctx, SQL, hotel.Id)
	helper.PanicIfError(err)
}

func (repository HotelRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, hotelId int) (domain.Hotel, error) {
	SQL := "select id, name, address, city, zip_code, rate from hotels where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, hotelId)
	helper.PanicIfError(err)
	defer rows.Close()

	hotel := domain.Hotel{}
	if rows.Next() {
		err := rows.Scan(&hotel.Id, &hotel.Name, &hotel.Address, &hotel.City, &hotel.ZipCode, &hotel.Rate)
		helper.PanicIfError(err)
		return hotel, nil
	} else {
		return hotel, errors.New("hotel is not found")
	}
}

func (repository HotelRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Hotel {
	SQL := "select id, name, address, city, zip_code, rate from hotels"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var hotels []domain.Hotel
	for rows.Next() {
		hotel := domain.Hotel{}
		err := rows.Scan(&hotel.Id, &hotel.Name, &hotel.Address, &hotel.City, &hotel.ZipCode, &hotel.Rate)
		helper.PanicIfError(err)
		hotels = append(hotels, hotel)
	}
	return hotels
}

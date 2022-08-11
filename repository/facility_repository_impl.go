package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type FacilityRepositoryImpl struct {
}

func NewFacilityRepository() FacilityRepository {
	return &FacilityRepositoryImpl{}
}

func (repository FacilityRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, facility domain.Facility) domain.Facility {
	SQL := "insert into facilities(name, description) values (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, facility.Name, facility.Description)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	facility.Id = int(id)
	return facility
}

func (repository FacilityRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, facility domain.Facility) domain.Facility {
	SQL := "update facilities set name = ?, description = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, facility.Name, facility.Description, facility.Id)
	helper.PanicIfError(err)

	return facility
}

func (repository FacilityRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, facility domain.Facility) {
	SQL := "delete from facilities where id = ?"
	_, err := tx.ExecContext(ctx, SQL, facility.Id)
	helper.PanicIfError(err)
}

func (repository FacilityRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, facilityId int) (domain.Facility, error) {
	SQL := "select id, name, description from facilities where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, facilityId)
	helper.PanicIfError(err)
	defer rows.Close()

	facility := domain.Facility{}
	if rows.Next() {
		err := rows.Scan(&facility.Id, &facility.Name, &facility.Description)
		helper.PanicIfError(err)
		return facility, nil
	} else {
		return facility, errors.New("Id is not found")
	}
}

func (repository FacilityRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Facility {
	SQL := "select id, name, description from facilities"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var facilities []domain.Facility
	for rows.Next() {
		facility := domain.Facility{}
		err := rows.Scan(&facility.Id, &facility.Name, &facility.Description)
		helper.PanicIfError(err)
		facilities = append(facilities, facility)
	}
	return facilities
}

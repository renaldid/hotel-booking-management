package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type EmployeeRepositoryImpl struct {
}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{}
}

func (repository EmployeeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee {
	SQL := "insert into employees(role_id, hotel_id, name, gender, address, email, phone_number) values (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, employee.RoleId, employee.HotelId, employee.Name, employee.Gender, employee.Address, employee.Email, employee.PhoneNumber)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	employee.Id = int(id)
	return employee
}

func (repository EmployeeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee {
	SQL := "update employees set hotel_id = ?, name = ?, gender = ?, address = ?, email = ?, phone_number = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, employee.HotelId, employee.Name, employee.Gender, employee.Address, employee.Email, employee.PhoneNumber, employee.Id)
	helper.PanicIfError(err)

	return employee
}

func (repository EmployeeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, employee domain.Employee) {
	SQL := "delete from employees where id = ?"
	_, err := tx.ExecContext(ctx, SQL, employee.Id)
	helper.PanicIfError(err)
}

func (repository EmployeeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, employeeId int) (domain.Employee, error) {
	SQL := "select e.id, e.role_id, e.hotel_id, e.name, e.gender, e.address, e.email, e.phone_number, r.role_name, h.name as hotel_name from employees e inner join roles r on e.role_id=r.id inner join hotels h on e.hotel_id=h.id where e.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, employeeId)
	helper.PanicIfError(err)
	defer rows.Close()

	employee := domain.Employee{}
	if rows.Next() {
		err := rows.Scan(&employee.Id, &employee.RoleId, &employee.HotelId, &employee.Name, &employee.Gender, &employee.Address, &employee.Email, &employee.PhoneNumber, &employee.Role, &employee.Hotel)
		helper.PanicIfError(err)
		return employee, nil
	} else {
		return employee, errors.New("employee is not found")
	}
}

func (repository EmployeeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Employee {
	SQL := "select e.id, e.role_id, e.hotel_id, e.name, e.gender, e.address, e.email, e.phone_number, r.role_name, h.name as hotel_name from employees e inner join roles r on e.role_id=r.id inner join hotels h on e.hotel_id=h.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var employees []domain.Employee
	for rows.Next() {
		employee := domain.Employee{}
		err := rows.Scan(&employee.Id, &employee.RoleId, &employee.HotelId, &employee.Name, &employee.Gender, &employee.Address, &employee.Email, &employee.PhoneNumber, &employee.Role, &employee.Hotel)
		helper.PanicIfError(err)
		employees = append(employees, employee)
	}
	return employees
}

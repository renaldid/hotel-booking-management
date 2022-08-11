package service

import (
	"context"
	"database/sql"
	"github.com/renaldid/hotel_booking_management.git/exception"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
	"github.com/renaldid/hotel_booking_management.git/model/web"
	"github.com/renaldid/hotel_booking_management.git/repository"

	"github.com/go-playground/validator/v10"
)

type EmployeeServiceImpl struct {
	EmployeeRepository repository.EmployeeRepository
	DB                 *sql.DB
	validate           *validator.Validate
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository, DB *sql.DB, validate *validator.Validate) EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
		DB:                 DB,
		validate:           validate,
	}
}

func (service *EmployeeServiceImpl) Create(ctx context.Context, request web.EmployeeCreateRequest) web.EmployeeResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee := domain.Employee{
		RoleId:      request.RoleId,
		HotelId:     request.HotelId,
		Name:        request.Name,
		Gender:      request.Gender,
		Address:     request.Address,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}

	employee = service.EmployeeRepository.Create(ctx, tx, employee)
	return helper.ToEmployeeResponse(employee)
}

func (service *EmployeeServiceImpl) Update(ctx context.Context, request web.EmployeeUpdateRequest) web.EmployeeResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := service.EmployeeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	employee.RoleId = request.RoleId
	employee.HotelId = request.HotelId
	employee.Name = request.Name
	employee.Gender = request.Gender
	employee.Address = request.Address
	employee.Email = request.Email
	employee.PhoneNumber = request.PhoneNumber

	employee = service.EmployeeRepository.Update(ctx, tx, employee)

	return helper.ToEmployeeResponse(employee)
}

func (service *EmployeeServiceImpl) Delete(ctx context.Context, employeeId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := service.EmployeeRepository.FindById(ctx, tx, employeeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.EmployeeRepository.Delete(ctx, tx, employee)
}

func (service *EmployeeServiceImpl) FindById(ctx context.Context, employeeId int) web.EmployeeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := service.EmployeeRepository.FindById(ctx, tx, employeeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToEmployeeResponse(employee)
}

func (service *EmployeeServiceImpl) FindAll(ctx context.Context) []web.EmployeeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employees := service.EmployeeRepository.FindAll(ctx, tx)

	return helper.ToEmployeeResponses(employees)
}

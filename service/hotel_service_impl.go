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

type HotelServiceImpl struct {
	HotelRepository repository.HotelRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewHotelService(hotelRepository repository.HotelRepository, DB *sql.DB, validate *validator.Validate) HotelService {
	return &HotelServiceImpl{
		HotelRepository: hotelRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *HotelServiceImpl) Create(ctx context.Context, request web.HotelCreateRequest) web.HotelResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hotel := domain.Hotel{
		Name:    request.Name,
		Address: request.Address,
		City:    request.City,
		ZipCode: request.ZipCode,
		Rate:    request.Rate,
	}

	hotel = service.HotelRepository.Create(ctx, tx, hotel)
	return helper.ToHotelResponse(hotel)
}

func (service *HotelServiceImpl) Update(ctx context.Context, request web.HotelUpdateRequest) web.HotelResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hotel, err := service.HotelRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	hotel.Id = request.Id
	hotel.Name = request.Name
	hotel.Address = request.Address
	hotel.ZipCode = request.ZipCode
	hotel.Rate = request.Rate

	hotel = service.HotelRepository.Update(ctx, tx, hotel)
	return helper.ToHotelResponse(hotel)
}

func (service *HotelServiceImpl) Delete(ctx context.Context, hotelId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hotel, err := service.HotelRepository.FindById(ctx, tx, hotelId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.HotelRepository.Delete(ctx, tx, hotel)
}

func (service *HotelServiceImpl) FindById(ctx context.Context, hotelId int) web.HotelResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hotel, err := service.HotelRepository.FindById(ctx, tx, hotelId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToHotelResponse(hotel)
}

func (service *HotelServiceImpl) FindAll(ctx context.Context) []web.HotelResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hotels := service.HotelRepository.FindAll(ctx, tx)

	return helper.ToHotelResponses(hotels)
}

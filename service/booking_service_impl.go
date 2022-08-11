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

type BookingServiceImpl struct {
	bookingRepository repository.BookingRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewBookingService(bookingRepository repository.BookingRepository, DB *sql.DB, Validate *validator.Validate) BookingService {
	return &BookingServiceImpl{
		bookingRepository: bookingRepository,
		DB:                DB,
		Validate:          Validate,
	}
}

func (service *BookingServiceImpl) Create(ctx context.Context, request web.BookingCreateRequest) web.BookingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	booking := domain.Booking{
		MeetingRoomId: request.MeetingRoomId,
		InvoiceId:     request.InvoiceId,
		Status:        request.Status,
	}
	booking = service.bookingRepository.Create(ctx, tx, booking)
	return helper.ToBookingResponse(booking)
}

func (service *BookingServiceImpl) Update(ctx context.Context, request web.BookingUpdateRequest) web.BookingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	booking, err := service.bookingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	booking.Id = request.Id
	booking.MeetingRoomId = request.MeetingRoomId
	booking.InvoiceId = request.InvoiceId
	booking.Status = request.Status

	booking = service.bookingRepository.Update(ctx, tx, booking)

	return helper.ToBookingResponse(booking)
}

func (service *BookingServiceImpl) Delete(ctx context.Context, bookingId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	booking, err := service.bookingRepository.FindById(ctx, tx, bookingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.bookingRepository.Delete(ctx, tx, booking)
}

func (service *BookingServiceImpl) FindById(ctx context.Context, bookingId int) web.BookingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	booking, err := service.bookingRepository.FindById(ctx, tx, bookingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToBookingResponse(booking)
}

func (service *BookingServiceImpl) FindAll(ctx context.Context) []web.BookingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	bookings := service.bookingRepository.FindAll(ctx, tx)

	return helper.ToBookingResponses(bookings)
}

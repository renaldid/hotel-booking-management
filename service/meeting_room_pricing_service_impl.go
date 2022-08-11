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

type MeetingRoomPricingServiceImpl struct {
	MeetingRoomPricingRepository repository.MeetingRoomPricingRepository
	DB                           *sql.DB
	Validate                     *validator.Validate
}

func NewMeetingRoomPricingService(meetingRoomPricingRepository repository.MeetingRoomPricingRepository, DB *sql.DB, validate *validator.Validate) MeetingRoomPricingService {
	return &MeetingRoomPricingServiceImpl{
		MeetingRoomPricingRepository: meetingRoomPricingRepository,
		DB:                           DB,
		Validate:                     validate,
	}
}

func (service *MeetingRoomPricingServiceImpl) Create(ctx context.Context, request web.MeetingRoomPricingCreateRequest) web.MeetingRoomPricingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoomPricing := domain.MeetingRoomPricing{
		MeetingRoomId: request.MeetingRoomId,
		Price:         request.Price,
		PriceType:     request.PriceType,
	}

	meetingRoomPricing = service.MeetingRoomPricingRepository.Create(ctx, tx, meetingRoomPricing)

	return helper.ToMeetingRoomPricingResponse(meetingRoomPricing)
}

func (service *MeetingRoomPricingServiceImpl) Update(ctx context.Context, request web.MeetingRoomPricingUpdateRequest) web.MeetingRoomPricingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoomPricing, err := service.MeetingRoomPricingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	meetingRoomPricing.MeetingRoomId = request.MeetingRoomId
	meetingRoomPricing.Price = request.Price
	meetingRoomPricing.PriceType = request.PriceType

	meetingRoomPricing = service.MeetingRoomPricingRepository.Update(ctx, tx, meetingRoomPricing)

	return helper.ToMeetingRoomPricingResponse(meetingRoomPricing)
}

func (service *MeetingRoomPricingServiceImpl) Delete(ctx context.Context, meetingRoomPricingId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoomPricing, err := service.MeetingRoomPricingRepository.FindById(ctx, tx, meetingRoomPricingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MeetingRoomPricingRepository.Delete(ctx, tx, meetingRoomPricing)
}

func (service *MeetingRoomPricingServiceImpl) FindById(ctx context.Context, meetingRoomPricingId int) web.MeetingRoomPricingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoomPricing, err := service.MeetingRoomPricingRepository.FindById(ctx, tx, meetingRoomPricingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMeetingRoomPricingResponse(meetingRoomPricing)
}

func (service *MeetingRoomPricingServiceImpl) FindAll(ctx context.Context) []web.MeetingRoomPricingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoomPricings := service.MeetingRoomPricingRepository.FindAll(ctx, tx)

	return helper.ToMeetingRoomPricingResponses(meetingRoomPricings)
}

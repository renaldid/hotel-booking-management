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

type MeetingRoomServiceImpl struct {
	meetingRoomRepository repository.MeetingRoomRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewMeetingRoomService(meetingRoomRepository repository.MeetingRoomRepository, DB *sql.DB, validate *validator.Validate) MeetingRoomService {
	return &MeetingRoomServiceImpl{
		meetingRoomRepository: meetingRoomRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *MeetingRoomServiceImpl) Create(ctx context.Context, request web.MeetingRoomCreateRequest) web.MeetingRoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoom := domain.MeetingRoom{
		FloorId:    request.FloorId,
		Name:       request.Name,
		Capacity:   request.Capacity,
		FacilityId: request.FacilityId,
	}
	meetingRoom = service.meetingRoomRepository.Create(ctx, tx, meetingRoom)
	return helper.ToMeetingRoomResponse(meetingRoom)
}

func (service *MeetingRoomServiceImpl) Update(ctx context.Context, request web.MeetingRoomUpdateRequest) web.MeetingRoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoom, err := service.meetingRoomRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	meetingRoom.Id = request.Id
	meetingRoom.FloorId = request.FloorId
	meetingRoom.Name = request.Name
	meetingRoom.Capacity = request.Capacity
	meetingRoom.FacilityId = request.FacilityId

	meetingRoom = service.meetingRoomRepository.Update(ctx, tx, meetingRoom)

	return helper.ToMeetingRoomResponse(meetingRoom)
}

func (service *MeetingRoomServiceImpl) Delete(ctx context.Context, meetingRoomId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoom, err := service.meetingRoomRepository.FindById(ctx, tx, meetingRoomId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.meetingRoomRepository.Delete(ctx, tx, meetingRoom)
}

func (service *MeetingRoomServiceImpl) FindById(ctx context.Context, meetingRoomId int) web.MeetingRoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoom, err := service.meetingRoomRepository.FindById(ctx, tx, meetingRoomId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMeetingRoomResponse(meetingRoom)
}

func (service *MeetingRoomServiceImpl) FindAll(ctx context.Context) []web.MeetingRoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRooms := service.meetingRoomRepository.FindAll(ctx, tx)

	return helper.ToMeetingRoomResponses(meetingRooms)
}

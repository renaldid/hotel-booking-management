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

type FacilityServiceImpl struct {
	FacilityRepository repository.FacilityRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewFacilityService(facilityRepository repository.FacilityRepository, DB *sql.DB, validate *validator.Validate) FacilityService {
	return &FacilityServiceImpl{
		FacilityRepository: facilityRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *FacilityServiceImpl) Create(ctx context.Context, request web.FacilityCreateRequest) web.FacilityResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	facility := domain.Facility{
		Name:        request.Name,
		Description: request.Description,
	}

	facility = service.FacilityRepository.Create(ctx, tx, facility)

	return helper.ToFacilityResponse(facility)
}

func (service *FacilityServiceImpl) Update(ctx context.Context, request web.FacilityUpdateRequest) web.FacilityResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	facility, err := service.FacilityRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	facility.Name = request.Name
	facility.Description = request.Description

	facility = service.FacilityRepository.Update(ctx, tx, facility)

	return helper.ToFacilityResponse(facility)
}

func (service *FacilityServiceImpl) Delete(ctx context.Context, facilityId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	facility, err := service.FacilityRepository.FindById(ctx, tx, facilityId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.FacilityRepository.Delete(ctx, tx, facility)
}

func (service *FacilityServiceImpl) FindById(ctx context.Context, facilityId int) web.FacilityResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	facility, err := service.FacilityRepository.FindById(ctx, tx, facilityId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToFacilityResponse(facility)
}

func (service *FacilityServiceImpl) FindAll(ctx context.Context) []web.FacilityResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	facilities := service.FacilityRepository.FindAll(ctx, tx)

	return helper.ToFacilityResponses(facilities)
}

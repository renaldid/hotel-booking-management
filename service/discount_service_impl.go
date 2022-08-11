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
	"github.com/sirupsen/logrus"
)

type DiscountServiceImpl struct {
	DiscountRepository repository.DiscountRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewDiscountService(discountRepository repository.DiscountRepository, DB *sql.DB, validate *validator.Validate) DiscountService {
	return &DiscountServiceImpl{
		DiscountRepository: discountRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *DiscountServiceImpl) Create(ctx context.Context, request web.DiscountCreateRequest) web.DiscountResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discount := domain.Discount{
		EmployeeId:    request.EmployeeId,
		HotelId:       request.HotelId,
		MeetingRoomId: request.MeetingRoomId,
		Rate:          request.Rate,
		Status:        request.Status,
		RequestDate:   request.RequestDate,
	}

	discount = service.DiscountRepository.Create(ctx, tx, discount)
	return helper.ToDiscountResponse(discount)
}

func (service *DiscountServiceImpl) Update(ctx context.Context, request web.DiscountUpdateRequest) web.DiscountResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discount, err := service.DiscountRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	discount.EmployeeId = request.EmployeeId
	discount.HotelId = request.HotelId
	discount.MeetingRoomId = request.MeetingRoomId
	discount.Rate = request.Rate
	discount.Status = request.Status
	discount.RequestDate = request.RequestDate

	discount = service.DiscountRepository.Update(ctx, tx, discount)

	return helper.ToDiscountResponse(discount)

}

func (service *DiscountServiceImpl) Delete(ctx context.Context, discountId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discount, err := service.DiscountRepository.FindById(ctx, tx, discountId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.DiscountRepository.Delete(ctx, tx, discount)
}

func (service *DiscountServiceImpl) FindById(ctx context.Context, discountId int) web.DiscountResponse {
	logrus.Info("Discount ser Find By Id star")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discount, err := service.DiscountRepository.FindById(ctx, tx, discountId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	logrus.Info("Discount ser Find By Id end")
	return helper.ToDiscountResponse(discount)
}

func (service *DiscountServiceImpl) FindAll(ctx context.Context) []web.DiscountResponse {
	logrus.Info("Discount ser Find al star")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discounts := service.DiscountRepository.FindAll(ctx, tx)
	logrus.Info("Discount ser Find al end")
	return helper.ToDiscountResponses(discounts)
}

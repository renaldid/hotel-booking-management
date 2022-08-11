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

type InvoiceServiceImpl struct {
	invoiceRepository repository.InvoiceRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewInvoiceService(invoiceRepository repository.InvoiceRepository, DB *sql.DB, Validate *validator.Validate) InvoiceService {
	return &InvoiceServiceImpl{
		invoiceRepository: invoiceRepository,
		DB:                DB,
		Validate:          Validate,
	}
}

func (service *InvoiceServiceImpl) Create(ctx context.Context, request web.InvoiceCreateRequest) web.InvoiceResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	invoice := domain.Invoice{
		Number:               request.Number,
		EmpeloyeeId:          request.EmpeloyeeId,
		MeetingRoomPricingId: request.MeetingRoomPricingId,
		DiscountId:           request.DiscountId,
		Pic:                  request.Pic,
	}
	invoice = service.invoiceRepository.Create(ctx, tx, invoice)
	return helper.ToInvoiceResponse(invoice)
}

func (service *InvoiceServiceImpl) Update(ctx context.Context, request web.InvoiceUpdateRequest) web.InvoiceResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	invoice, err := service.invoiceRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	invoice.Id = request.Id
	invoice.Number = request.Number
	invoice.EmpeloyeeId = request.EmpeloyeeId
	invoice.MeetingRoomPricingId = request.MeetingRoomPricingId
	invoice.DiscountId = request.DiscountId
	invoice.Pic = request.Pic

	invoice = service.invoiceRepository.Update(ctx, tx, invoice)

	return helper.ToInvoiceResponse(invoice)
}

func (service *InvoiceServiceImpl) Delete(ctx context.Context, invoiceId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	invoice, err := service.invoiceRepository.FindById(ctx, tx, invoiceId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.invoiceRepository.Delete(ctx, tx, invoice)
}

func (service *InvoiceServiceImpl) FindById(ctx context.Context, invoiceId int) web.InvoiceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	invoice, err := service.invoiceRepository.FindById(ctx, tx, invoiceId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToInvoiceResponse(invoice)
}

func (service *InvoiceServiceImpl) FindAll(ctx context.Context) []web.InvoiceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	invoices := service.invoiceRepository.FindAll(ctx, tx)

	return helper.ToInvoiceResponses(invoices)
}

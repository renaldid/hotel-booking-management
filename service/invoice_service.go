package service

import (
	"context"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

type InvoiceService interface {
	Create(ctx context.Context, request web.InvoiceCreateRequest) web.InvoiceResponse
	Update(ctx context.Context, request web.InvoiceUpdateRequest) web.InvoiceResponse
	Delete(ctx context.Context, invoiceId int)
	FindById(ctx context.Context, invoiceId int) web.InvoiceResponse
	FindAll(ctx context.Context) []web.InvoiceResponse
}

package repository

import (
	"context"
	"database/sql"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type InvoiceRepository interface {
	Create(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) domain.Invoice
	Update(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) domain.Invoice
	Delete(ctx context.Context, tx *sql.Tx, invoice domain.Invoice)
	FindById(ctx context.Context, tx *sql.Tx, invoiceId int) (domain.Invoice, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Invoice
}

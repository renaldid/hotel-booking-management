package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type InvoiceRepositoryImpl struct {
}

func NewInvoiceRepository() InvoiceRepository {
	return &InvoiceRepositoryImpl{}
}

func (repository InvoiceRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) domain.Invoice {
	SQL := "insert into invoices(number, employee_id, meeting_room_pricing_id, discount_id, pic) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, invoice.Number, invoice.EmpeloyeeId, invoice.MeetingRoomPricingId, invoice.DiscountId, invoice.Pic)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	invoice.Id = int(id)
	return invoice
}

func (repository InvoiceRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) domain.Invoice {
	SQL := "update invoices set number = ?, employee_id = ?, meeting_room_pricing_id = ?, discount_id = ?, pic = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, invoice.Number, invoice.EmpeloyeeId, invoice.MeetingRoomPricingId, invoice.DiscountId, invoice.Pic, invoice.Id)
	helper.PanicIfError(err)

	return invoice
}

func (repository InvoiceRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) {
	SQL := "delete from invoices where id = ?"
	_, err := tx.ExecContext(ctx, SQL, invoice.Id)
	helper.PanicIfError(err)
}

func (repository InvoiceRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, invoiceId int) (domain.Invoice, error) {
	SQL := "select i.id, i.number, i.employee_id, i.meeting_room_pricing_id, i.discount_id, i.pic, e.name as employee_name, m.price as price, m.price_type as price_type, d.rate as discount_rate, d.status as discount_status " +
		"from (((invoices i inner join employees e on i.employee_id=e.id) inner join meeting_room_pricings m on i.meeting_room_pricing_id=m.id) inner join discounts d on i.discount_id=d.id) where i.id=?"
	rows, err := tx.QueryContext(ctx, SQL, invoiceId)
	helper.PanicIfError(err)
	defer rows.Close()

	invoice := domain.Invoice{}
	if rows.Next() {
		err := rows.Scan(&invoice.Id, &invoice.Number, &invoice.EmpeloyeeId, &invoice.MeetingRoomPricingId, &invoice.DiscountId, &invoice.Pic, &invoice.EmpeloyeeName, &invoice.Price, &invoice.PriceType, &invoice.DiscountRate, &invoice.DiscountStatus)
		helper.PanicIfError(err)
		return invoice, nil
	} else {
		return invoice, errors.New("invoice is not found")
	}
}

func (repository InvoiceRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Invoice {
	SQL := "select i.id, i.number, i.employee_id, i.meeting_room_pricing_id, i.discount_id, i.pic, e.name as empeloyee_name, m.price as price, m.price_type as price_type, d.rate as discount_rate, d.status as discount_status from (((invoices i inner join employees e on i.employee_id=e.id) inner join meeting_room_pricings m on i.meeting_room_pricing_id=m.id) inner join discounts d on i.discount_id=d.id)"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var invoices []domain.Invoice
	for rows.Next() {
		invoice := domain.Invoice{}
		err := rows.Scan(&invoice.Id, &invoice.Number, &invoice.EmpeloyeeId, &invoice.MeetingRoomPricingId, &invoice.DiscountId, &invoice.Pic, &invoice.EmpeloyeeName, &invoice.Price, &invoice.PriceType, &invoice.DiscountRate, &invoice.DiscountStatus)
		helper.PanicIfError(err)
		invoices = append(invoices, invoice)
	}
	return invoices
}

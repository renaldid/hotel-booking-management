package web

type InvoiceCreateRequest struct {
	Number               int    `validate:"required" json:"number"`
	EmpeloyeeId          int    `validate:"required" json:"employee_id"`
	MeetingRoomPricingId int    `validate:"required" json:"meeting_room_pricing_id"`
	DiscountId           int    `validate:"required" json:"discount_id"`
	Pic                  string `validate:"required" json:"pic"`
}

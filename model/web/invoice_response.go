package web

type InvoiceResponse struct {
	Id                   int    `json:"id"`
	Number               int    `json:"number"`
	EmpeloyeeId          int    `json:"employee_id"`
	MeetingRoomPricingId int    `json:"meeting_room_pricing_id"`
	DiscountId           int    `json:"discount_id"`
	Pic                  string `json:"pic"`
	EmpeloyeeName        string `json:"empeloyee_name"`
	Price                string `json:"price"`
	PriceType            string `json:"price_type"`
	DiscountRate         string `json:"discount_rate"`
	DiscountStatus       string `json:"discount_status"`
}

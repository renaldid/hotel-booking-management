package domain

type Invoice struct {
	Id                   int
	Number               int
	EmpeloyeeId          int
	MeetingRoomPricingId int
	DiscountId           int
	Pic                  string
	EmpeloyeeName        string
	Price                string
	PriceType            string
	DiscountRate         string
	DiscountStatus       string
}

package web

type MeetingRoomPricingCreateRequest struct {
	MeetingRoomId int    `validate:"required" json:"meeting_room_id"`
	Price         string `validate:"required,min=1,max=20" json:"price"`
	PriceType     string `validate:"required,min=1,max=100" json:"price_type"`
}

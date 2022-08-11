package web

type MeetingRoomPricingUpdateRequest struct {
	Id            int    `validate:"required"`
	MeetingRoomId int    `validate:"required" json:"meeting_room_id"`
	Price         string `validate:"required,max=20,min=1" json:"price"`
	PriceType     string `validate:"required,max=100,min=1" json:"price_type"`
}

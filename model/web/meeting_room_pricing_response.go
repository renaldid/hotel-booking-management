package web

type MeetingRoomPricingResponse struct {
	Id            int    `json:"id"`
	MeetingRoomId int    `json:"meeting_room_id"`
	Price         string `json:"price"`
	PriceType     string `json:"price_type"`
	Name          string `json:"name"`
	Capacity      string `json:"capacity"`
}

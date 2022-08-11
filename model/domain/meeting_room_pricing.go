package domain

type MeetingRoomPricing struct {
	Id            int
	MeetingRoomId int
	Price         string
	PriceType     string
	Name          string
	Capacity      string
}

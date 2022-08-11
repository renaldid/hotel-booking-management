package web

type BookingUpdateRequest struct {
	Id            int    `validate:"required"`
	MeetingRoomId int    `validate:"required" json:"meeting_room_id"`
	InvoiceId     int    `validate:"required" json:"invoice_id"`
	Status        string `validate:"required,min=1,max=100" json:"status"`
}

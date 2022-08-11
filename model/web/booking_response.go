package web

type BookingResponse struct {
	Id              int    `json:"id"`
	MeetingRoomId   int    `json:"meeting_room_id"`
	InvoiceId       int    `json:"invoice_id"`
	Status          string `json:"status"`
	MeetingRoomName string `json:"meeting_room_name"`
	InvoiceNumber   int    `json:"invoice_number"`
	PicName         string `json:"pic_name"`
}

package domain

type Booking struct {
	Id              int
	MeetingRoomId   int
	InvoiceId       int
	Status          string
	MeetingRoomName string
	InvoiceNumber   int
	PicName         string
}

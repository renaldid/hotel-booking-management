package web

type DiscountCreateRequest struct {
	EmployeeId    int    `validate:"required"json:"employee_id"`
	HotelId       int    `validate:"required"json:"hotel_id"`
	MeetingRoomId int    `validate:"required"json:"meeting_room_id"`
	Rate          string `validate:"required,min=1,max=10"json:"rate"`
	Status        string `validate:"required,min=1,max=10"json:"status"`
	RequestDate   string `validate:"required,min=1,max=100"json:"request_date"`
}

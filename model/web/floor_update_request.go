package web

type FloorUpdateRequest struct {
	Id          int    `validate:"required"`
	HotelId     int    `validate:"required" json:"hotel_id"`
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required,min=1,max=100" json:"description"`
}

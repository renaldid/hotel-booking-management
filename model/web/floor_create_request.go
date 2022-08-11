package web

type FloorCreateRequest struct {
	HotelId     int    `validate:"required"json:"hotel_id"`
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required,min=1,max=100" json:"description"`
}

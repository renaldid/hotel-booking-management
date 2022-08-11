package web

type FloorResponse struct {
	Id          int    `json:"id"`
	HotelId     int    `json:"hotel_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	HotelName   string `json:"hotel_name"`
}

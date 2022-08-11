package web

type HotelResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
	Rate    string `json:"rate"`
}

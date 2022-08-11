package web

type HotelUpdateRequest struct {
	Id      int    `validate:"required"`
	Name    string `validate:"min=1,max=100" json:"name"`
	Address string `validate:"min=1,max=100" json:"address"`
	City    string `validate:"min=1,max=100" json:"city"`
	ZipCode string `validate:"min=1,max=100" json:"zip_code"`
	Rate    string `validate:"min=1,max=100" json:"rate"`
}

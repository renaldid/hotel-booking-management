package web

type HotelCreateRequest struct {
	Name    string `validate:"required,min=1,max=100" json:"name"`
	Address string `validate:"required,min=10,max=100" json:"address"`
	City    string `validate:"required,min=5,max=100" json:"city"`
	ZipCode string `validate:"required,min=1,max=100" json:"zip_code"`
	Rate    string `validate:"required,min=1,max=100"json:"rate"`
}

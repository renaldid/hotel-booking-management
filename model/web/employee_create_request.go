package web

type EmployeeCreateRequest struct {
	RoleId      int    `validate:"required" json:"role_id"`
	HotelId     int    `validate:"required" json:"hotel_id"`
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Gender      string `validate:"required,min=1,max=100" json:"gender"`
	Address     string `validate:"required,min=1,max=250" json:"address"`
	Email       string `validate:"required,min=8,max=100" json:"email"`
	PhoneNumber string `validate:"required,min=1,max=100" json:"phone_number"`
}

package web

type EmployeeResponse struct {
	Id          int    `json:"id"`
	RoleId      int    `json:"role_id"`
	HotelId     int    `json:"hotel_id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	Hotel       string `json:"hotel"`
}

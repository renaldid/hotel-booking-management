package helper

import (
	"github.com/renaldid/hotel_booking_management.git/model/domain"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

func ToEmployeeResponse(employee domain.Employee) web.EmployeeResponse {

	return web.EmployeeResponse{
		Id:          employee.Id,
		RoleId:      employee.RoleId,
		HotelId:     employee.HotelId,
		Name:        employee.Name,
		Gender:      employee.Gender,
		Address:     employee.Address,
		Email:       employee.Email,
		PhoneNumber: employee.PhoneNumber,
		Role:        employee.Role,
		Hotel:       employee.Hotel,
	}
}

func ToEmployeeResponses(employees []domain.Employee) []web.EmployeeResponse {
	var employeeResponses []web.EmployeeResponse
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, ToEmployeeResponse(employee))
	}
	return employeeResponses
}

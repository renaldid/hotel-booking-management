package helper

import (
	"github.com/renaldid/hotel_booking_management.git/model/domain"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

func ToDiscountResponse(discount domain.Discount) web.DiscountResponse {
	return web.DiscountResponse{
		Id:            discount.Id,
		EmployeeId:    discount.EmployeeId,
		HotelId:       discount.HotelId,
		MeetingRoomId: discount.MeetingRoomId,
		Rate:          discount.Rate,
		Status:        discount.Status,
		RequestDate:   discount.RequestDate,
		EmployeeName:  discount.EmployeeName,
		HotelName:     discount.HotelName,
		RoomName:      discount.RoomName,
	}
}

func ToDiscountResponses(discounts []domain.Discount) []web.DiscountResponse {
	var discountResponses []web.DiscountResponse
	for _, discount := range discounts {
		discountResponses = append(discountResponses, ToDiscountResponse(discount))
	}
	return discountResponses
}

package helper

import (
	"github.com/renaldid/hotel_booking_management.git/model/domain"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

func ToFloorResponse(floor domain.Floor) web.FloorResponse {
	return web.FloorResponse{
		Id:          floor.Id,
		HotelId:     floor.HotelId,
		Name:        floor.Name,
		Description: floor.Description,
		HotelName:   floor.HotelName,
	}
}

func ToFloorResponses(floors []domain.Floor) []web.FloorResponse {
	var floorResponses []web.FloorResponse
	for _, floor := range floors {
		floorResponses = append(floorResponses, ToFloorResponse(floor))
	}
	return floorResponses
}

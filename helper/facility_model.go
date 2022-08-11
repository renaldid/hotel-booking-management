package helper

import (
	"github.com/renaldid/hotel_booking_management.git/model/domain"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

func ToFacilityResponse(facility domain.Facility) web.FacilityResponse {
	return web.FacilityResponse{
		Id:          facility.Id,
		Name:        facility.Name,
		Description: facility.Description,
	}
}

func ToFacilityResponses(facilities []domain.Facility) []web.FacilityResponse {
	var facilityResponses []web.FacilityResponse
	for _, facility := range facilities {
		facilityResponses = append(facilityResponses, ToFacilityResponse(facility))
	}
	return facilityResponses
}

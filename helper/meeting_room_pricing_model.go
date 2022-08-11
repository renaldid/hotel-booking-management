package helper

import (
	"github.com/renaldid/hotel_booking_management.git/model/domain"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

func ToMeetingRoomPricingResponse(meetingRoomPricing domain.MeetingRoomPricing) web.MeetingRoomPricingResponse {
	return web.MeetingRoomPricingResponse{
		Id:            meetingRoomPricing.Id,
		MeetingRoomId: meetingRoomPricing.MeetingRoomId,
		Price:         meetingRoomPricing.Price,
		PriceType:     meetingRoomPricing.PriceType,
		Name:          meetingRoomPricing.Name,
		Capacity:      meetingRoomPricing.Capacity,
	}
}

func ToMeetingRoomPricingResponses(meetingRoomPricings []domain.MeetingRoomPricing) []web.MeetingRoomPricingResponse {
	var meetingRoomPricingResponses []web.MeetingRoomPricingResponse
	for _, meetingRoomPricing := range meetingRoomPricings {
		meetingRoomPricingResponses = append(meetingRoomPricingResponses, ToMeetingRoomPricingResponse(meetingRoomPricing))
	}
	return meetingRoomPricingResponses
}

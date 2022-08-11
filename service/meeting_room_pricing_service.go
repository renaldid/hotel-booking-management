package service

import (
	"context"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

type MeetingRoomPricingService interface {
	Create(ctx context.Context, request web.MeetingRoomPricingCreateRequest) web.MeetingRoomPricingResponse
	Update(ctx context.Context, request web.MeetingRoomPricingUpdateRequest) web.MeetingRoomPricingResponse
	Delete(ctx context.Context, meetingRoomPricingId int)
	FindById(ctx context.Context, meetingRoomPricingId int) web.MeetingRoomPricingResponse
	FindAll(ctx context.Context) []web.MeetingRoomPricingResponse
}

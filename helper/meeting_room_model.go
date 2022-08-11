package helper

import (
	"github.com/renaldid/hotel_booking_management.git/model/domain"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

func ToMeetingRoomResponse(meetingRoom domain.MeetingRoom) web.MeetingRoomResponse {
	return web.MeetingRoomResponse{
		Id:                  meetingRoom.Id,
		FloorId:             meetingRoom.FloorId,
		Name:                meetingRoom.Name,
		Capacity:            meetingRoom.Capacity,
		FacilityId:          meetingRoom.FacilityId,
		FloorName:           meetingRoom.FloorName,
		FloorDescription:    meetingRoom.FloorDescription,
		FacilityName:        meetingRoom.FacilityName,
		FacilityDescription: meetingRoom.FacilityDescription,
	}
}

func ToMeetingRoomResponses(meetingRooms []domain.MeetingRoom) []web.MeetingRoomResponse {
	var meetingRoomResponses []web.MeetingRoomResponse
	for _, meetingRoom := range meetingRooms {
		meetingRoomResponses = append(meetingRoomResponses, ToMeetingRoomResponse(meetingRoom))
	}
	return meetingRoomResponses
}

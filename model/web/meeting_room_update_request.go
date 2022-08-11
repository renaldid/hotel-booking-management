package web

type MeetingRoomUpdateRequest struct {
	Id         int    `validate:"required"`
	FloorId    int    `validate:"required" json:"floor_id"`
	Name       string `validate:"required,max=100,min=1" json:"name"`
	Capacity   string `validate:"required,max=100,min=1" json:"capacity"`
	FacilityId int    `validate:"required" json:"facility_id"`
}

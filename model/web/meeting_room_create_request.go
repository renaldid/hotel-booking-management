package web

type MeetingRoomCreateRequest struct {
	FloorId    int    `validate:"required" json:"floor_id"`
	Name       string `validate:"required,min=1,max=100" json:"name"`
	Capacity   string `validate:"required,min=1,max=100" json:"capacity"`
	FacilityId int    `validate:"required" json:"facility_id"`
}

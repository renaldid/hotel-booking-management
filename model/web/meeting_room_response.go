package web

type MeetingRoomResponse struct {
	Id                  int    `json:"id"`
	FloorId             int    `json:"floor_id"`
	Name                string `json:"name"`
	Capacity            string `json:"capacity"`
	FacilityId          int    `json:"facility_id"`
	FloorName           string `json:"floor_name"`
	FloorDescription    string `json:"floor_description"`
	FacilityName        string `json:"facility_name"`
	FacilityDescription string `json:"facility_description"`
}

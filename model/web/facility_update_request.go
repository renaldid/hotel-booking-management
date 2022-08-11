package web

type FacilityUpdateRequest struct {
	Id          int    `validate:"required"`
	Name        string `validate:"required,max=100,min=1" json:"name"`
	Description string `validate:"required,max=300,min=1" json:"description"`
}

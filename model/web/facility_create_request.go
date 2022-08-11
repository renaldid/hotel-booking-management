package web

type FacilityCreateRequest struct {
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required,min=1,max=300" json:"description"`
}

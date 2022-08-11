package web

type UserUpdateRequest struct {
	Id       int    `validate:"required"`
	Name     string `validate:"min=1,max=100" json:"name"`
	Email    string `validate:"min=10,max=100" json:"email"`
	Password string `validate:"min=5,max=100" json:"password"`
}

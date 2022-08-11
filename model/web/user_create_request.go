package web

type UserCreateRequest struct {
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Email    string `validate:"required,min=10,max=100" json:"email"`
	Password string `validate:"required,min=5,max=100" json:"password"`
	RoleId   int    `validate:"required,min=1,max=100" json:"role_id"`
}

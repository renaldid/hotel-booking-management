package web

type RoleCreateRequest struct {
	RoleName string `validate:"required,min=1,max=100" json:"role_name"`
}

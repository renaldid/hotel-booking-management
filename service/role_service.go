package service

import (
	"context"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

type RoleService interface {
	Create(ctx context.Context, request web.RoleCreateRequest) web.RoleResponse
	Update(ctx context.Context, request web.RoleUpdateRequest) web.RoleResponse
	Delete(ctx context.Context, roleId int)
	FindById(ctx context.Context, roleId int) web.RoleResponse
	FindAll(ctx context.Context) []web.RoleResponse
}

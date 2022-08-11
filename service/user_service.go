package service

import (
	"context"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}

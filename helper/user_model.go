package helper

import (
	"github.com/renaldid/hotel_booking_management.git/model/domain"
	"github.com/renaldid/hotel_booking_management.git/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {

	return web.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleId:   user.RoleId,
		RoleName: user.RoleName,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

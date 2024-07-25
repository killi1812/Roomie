package Helpers

import (
	"chatapp/server/dtos"
	"chatapp/server/models"
)

func MapNewUserDtoToUser(dto dtos.NewUserDto) models.User {
	return models.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

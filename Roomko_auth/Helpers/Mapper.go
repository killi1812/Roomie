package Helpers

import (
	"roomko/auth/dtos"
	"roomko/auth/models"
)

func MapNewUserDtoToUser(dto dtos.NewUserDto) models.User {
	return models.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

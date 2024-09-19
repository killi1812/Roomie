package models

import (
	"github.com/google/uuid"
	"roomko/auth/dtos"
)

type User struct {
	//TODO add guid
	Uuid     uuid.UUID `json:"uuid" bson:"uuid"`
	Username string    `json:"username" bson:"username"`
	Email    string    `json:"email" bson:"email"`
	Password string    `json:"password" bson:"password"`
}

func NewUser(dto dtos.NewUserDto, hashedPsswd string) User {
	return User{
		Uuid:     uuid.New(),
		Username: dto.Username,
		Email:    dto.Email,
		Password: hashedPsswd,
	}
}

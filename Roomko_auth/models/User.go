package models

import "github.com/google/uuid"

type User struct {
	//TODO add guid
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

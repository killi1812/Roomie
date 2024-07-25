package Services

import (
	"chatapp/server/Helpers"
	"chatapp/server/dtos"
	"chatapp/server/models"
	"encoding/json"
	"fmt"
	"os"
)

type UserService interface {
	CreateUser(dto dtos.NewUserDto)
}

func CreateUser(dto dtos.NewUserDto) (models.User, error) {
	file, err := os.Create("users.json")
	if err != nil {
		fmt.Println("Error creating a user")
		return models.User{}, err
	}
	defer file.Close()
	hashedPassword, err := Helpers.HashPassword(dto.Password)

	if err != nil {
		fmt.Println("Error hashing a password")
		return models.User{}, err
	}

	dto.Password = hashedPassword
	encoder := json.NewEncoder(file)
	err = encoder.Encode(dto)
	if err != nil {
		fmt.Println("Error encoding a user")
		return models.User{}, err
	}
	return Helpers.MapNewUserDtoToUser(dto), nil
}

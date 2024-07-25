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
	Login(dto dtos.UserAuthDto)
}

// TODO make a struct
const fileName = "users.json"

func CreateUser(dto dtos.NewUserDto) (models.User, error) {
	//TODO check if user already exists
	//TODO check if email is valid
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating a user")
		return models.User{}, err
	}
	defer file.Close()

	//TODO move to a better system
	var users []models.User
	json.NewDecoder(file).Decode(&users)
	for _, user := range users {
		if user.Username == dto.Username {
			return models.User{}, fmt.Errorf("User %s already exists", dto.Username)
		}
	}
	hashedPassword, err := Helpers.HashPassword(dto.Password)

	if err != nil {
		fmt.Println("Error hashing a password")
		return models.User{}, err
	}

	dto.Password = hashedPassword
	users = append(users, Helpers.MapNewUserDtoToUser(dto))
	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Println("Error encoding a user")
		return models.User{}, err
	}
	return Helpers.MapNewUserDtoToUser(dto), nil
}

func Login(dto dtos.UserAuthDto) (models.User, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening a file")
		return models.User{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var users []models.User
	err = decoder.Decode(&users)
	if err != nil {
		fmt.Println("Error decoding a file")
		return models.User{}, err
	}

	for _, user := range users {
		if user.Username == dto.Username {
			if Helpers.CheckPasswordHash(dto.Password, user.Password) {
				return user, nil
			}
		}
	}
	return models.User{}, fmt.Errorf("User %s not found", dto.Username)
}

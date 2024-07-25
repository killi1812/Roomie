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
	users, err := LoadUsers()
	for _, user := range users {
		if user.Username == dto.Username {
			return models.User{}, fmt.Errorf("User %s already exists", dto.Username)
		}
	}
	//TODO move to a better system
	if err != nil {
		fmt.Println("Error loading users")
		return models.User{}, err
	}
	hashedPassword, err := Helpers.HashPassword(dto.Password)
	if err != nil {
		fmt.Println("Error hashing a password")
		return models.User{}, err
	}
	file, err := os.Create(fileName)
	defer file.Close()

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
	users, err := LoadUsers()
	if err != nil {
		fmt.Println("Error loading users")
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

func LoadUsers() (users []models.User, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []models.User{}, nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		fmt.Println("Error decoding a file")
		return []models.User{}, err
	}
	return users, nil
}

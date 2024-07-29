package Services

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
	"roomko/auth/Helpers"
	"roomko/auth/dtos"
	"roomko/auth/models"
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
	user := Helpers.MapNewUserDtoToUser(dto)
	user.Uuid = uuid.New()
	users = append(users, user)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Println("Error encoding a user")
		return models.User{}, err
	}
	return Helpers.MapNewUserDtoToUser(dto), nil
}

func Login(dto dtos.UserAuthDto) (models.Certificate, error) {
	users, err := LoadUsers()
	if err != nil {
		fmt.Println("Error loading users")
		return models.Certificate{}, err
	}

	for _, user := range users {
		if user.Username == dto.Username {
			if Helpers.CheckPasswordHash(dto.Password, user.Password) {
				cert, err := GenerateCertificate(user)
				if err != nil {
					return models.Certificate{}, err
				}
				return cert, nil
			}
		}
	}
	return models.Certificate{}, fmt.Errorf("User %s not found", dto.Username)
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
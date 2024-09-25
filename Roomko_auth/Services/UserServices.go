package Services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"roomko/auth/Helpers"
	"roomko/auth/dtos"
	"roomko/auth/models"
)

type UserService interface {
	CreateUser(dto dtos.NewUserDto) (models.User, error)
	Login(dto dtos.UserAuthDto) (models.Certificate, error)
}

const fileName = "users.json"

type FileDb struct {
}

func (_ FileDb) CreateUser(dto dtos.NewUserDto) (models.User, error) {
	//TODO check if user already exists
	//TODO check if email is valid
	users, err := loadUsers()
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
	user := models.NewUser(dto, hashedPassword)
	users = append(users, user)

	file, err := os.Create(fileName)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Println("Error encoding a user")
		return models.User{}, err
	}
	return user, nil
}

func (_ FileDb) Login(dto dtos.UserAuthDto) (models.Certificate, error) {
	users, err := loadUsers()
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

func loadUsers() (users []models.User, err error) {
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

type MongoDb struct {
}

// TODO check if it needs to be an extenstion
func GetUser(username string) (user models.User, err error) {
	conn, closeConn := Helpers.GetConnect()
	defer closeConn()
	dbName := Helpers.GetConfig().DbName
	coll := conn.Database(dbName).Collection("Users")

	err = coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)

	if errors.Is(err, mongo.ErrNoDocuments) {
		//TODO remove after debugging
		fmt.Printf("no user")
		return models.User{}, err
	}

	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	return
}

func insertUser(user models.User) error {
	conn, closeConn := Helpers.GetConnect()
	defer closeConn()
	dbName := Helpers.GetConfig().DbName
	coll := conn.Database(dbName).Collection("Users")

	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (_ MongoDb) CreateUser(dto dtos.NewUserDto) (models.User, error) {
	_, err := GetUser(dto.Username)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return models.User{}, fmt.Errorf("User %s already exists", dto.Username)
	}
	hashedPassword, err := Helpers.HashPassword(dto.Password)
	if err != nil {
		fmt.Println("Error hashing a password")
		return models.User{}, err
	}
	user := models.NewUser(dto, hashedPassword)
	if err = insertUser(user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (_ MongoDb) Login(dto dtos.UserAuthDto) (models.Certificate, error) {
	user, err := GetUser(dto.Username)
	if err != nil {
		fmt.Println("Error retriving a user %s \n%s", dto.Username, err)
	}

	if user.Username == dto.Username {
		if Helpers.CheckPasswordHash(dto.Password, user.Password) {
			cert, err := GenerateCertificate(user)
			if err != nil {
				return models.Certificate{}, err
			}
			return cert, nil
		}
	}

	return models.Certificate{}, fmt.Errorf("User %s not found", dto.Username)
}

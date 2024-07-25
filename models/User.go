package models

type User struct {
	//TODO add guid
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

package routes

import (
	"chatapp/server/Services"
	"chatapp/server/dtos"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func AddRoutes(router *httprouter.Router) {
	baseRout := "api/v1/auth"
	router.GET(fmt.Sprintf("%s/ping", baseRout), ping)
	router.GET(fmt.Sprintf("%s/public-key", baseRout), getPublicKey)
	router.GET(fmt.Sprintf("%s/verify-certificate"), verifyCertificate)
	router.POST(fmt.Sprintf("%s/login"), login)
	router.POST(fmt.Sprintf("%s/register"), register)
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Pong")
}

func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var userAuth dtos.UserAuthDto
	err := json.NewDecoder(r.Body).Decode(&userAuth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cert, err := Services.Login(userAuth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"message":     fmt.Sprintf("User %s logged in successfully", cert.Username),
			"certificate": cert,
		})
}

func register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newUser dtos.NewUserDto
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := Services.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fmt.Sprintf("User %s created successfully", user.Username))
}

func getPublicKey(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//TODO implement
}

func verifyCertificate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//TODO implement
}

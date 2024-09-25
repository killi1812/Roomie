package main

import (
	"fmt"
	"net/http"
	"roomko/auth/Helpers"
	"roomko/auth/Services"
	auth "roomko/auth/routes"

	"github.com/julienschmidt/httprouter"
)

func main() {
	err := Services.GenerateKeyPair()
	if err != nil {
		fmt.Println("Error generating private and public keys\n", err)
		return
	}
	router := httprouter.New()
	auth.AuthAddRoutes(router)
	auth.PagesAddRoutes(router)
	config := Helpers.GetConfig()

	fmt.Println("Server is starting")
	err = http.ListenAndServeTLS(fmt.Sprintf(":%d", config.Port), "keys/Https_cert.pem", "keys/Https_key.pem", router)

	if err != nil {
		fmt.Println("Error starting server\n", err)
		fmt.Println("listening on http port", config.Port)
		http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router)
	}
}

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"roomko/auth/Services"
	auth "roomko/auth/routes"
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
	port := 8832
	err = http.ListenAndServeTLS(fmt.Sprintf(":%d", port), "keys/Https_cert.pem", "keys/Https_key.pem", router)

	if err != nil {
		fmt.Println("Error starting server\n", err)
		http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	}
}

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
	auth.AddRoutes(router)
	port := 8832
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

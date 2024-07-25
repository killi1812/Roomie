package main

import (
	"chatapp/server/Services"
	auth "chatapp/server/routes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	err := Services.GenerateKeyPair()
	if err != nil {
		fmt.Println("Error generating private and public keys\n %s", err)
		return
	}
	router := httprouter.New()
	auth.AddRoutes(router)
	port := 8832
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

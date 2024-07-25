package main

import (
	auth "chatapp/server/routes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//import _ "chatapp/server/docs"

func main() {
	router := httprouter.New()
	auth.AuthRoute(router)

	port := 8832
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

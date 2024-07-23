package auth

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateRoute() {
	router := httprouter.New()
	port := 8832

	router.GET("/auth/ping", ping)

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprintf(w, "Pong")
}

func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func register(w http.w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
  
}

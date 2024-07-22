package auth

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateRoute() {
	router := httprouter.New()
	port := 8832

	router.GET("/auth/ping/:word", ping)

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "%s", ps.ByName("name"))
}

func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

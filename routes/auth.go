package routes 

import (
	"fmt"
	"net/http"
  "encoding/json"
  "html/template"

	"github.com/julienschmidt/httprouter"
)

func AuthRoute() {
	router := httprouter.New()
	port := 8832

	router.GET("/auth/ping", ping)

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  json.NewEncoder(w).Encode("Pong")
  //  fmt.Fprintf(w, "Pong")
}

func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func register(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
  tmpl := template.Mustl(template.) 
  
}

package routes

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
)

func PagesAddRoutes(router *httprouter.Router) {
	baseRoute := "/gui/v1"

	router.GET("/", reroute)
	router.GET("/index.html", reroute)
	//TODO rename functions in auth.go
	router.GET(fmt.Sprintf("%s/login", baseRoute), loginPage)
}

func loginPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, filepath.Join("wwwroot", "Auth", "Login.html"))
}

func reroute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/gui/v1/index.html", http.StatusFound)
}

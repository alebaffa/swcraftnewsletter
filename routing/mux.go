package routing

import (
	"net/http"

	"github.com/alebaffa/newsletter-web/utils"
	"github.com/gorilla/mux"
)

func HomePageHandler(res http.ResponseWriter, req *http.Request) {
	utils.ExecuteTemplate("homepage", res)
}

func ContributeHandler(res http.ResponseWriter, req *http.Request) {
	utils.ExecuteTemplate("contribute", res)
}

func NewMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePageHandler)
	r.HandleFunc("/contribute", ContributeHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	http.Handle("/", r)
	return r
}

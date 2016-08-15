package routing

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePageHandler(res http.ResponseWriter, req *http.Request) {
	executeTemplate("homepage", res)
}

func executeTemplate(templateType string, res http.ResponseWriter) {
	templateType = templateType + ".html"
	templates := template.Must(template.ParseGlob("assets/*.html"))
	if err := templates.ExecuteTemplate(res, templateType, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func NewMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePageHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	http.Handle("/", r)
	return r
}

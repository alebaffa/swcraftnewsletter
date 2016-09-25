package routing

import (
	"net/http"

	"github.com/alebaffa/newsletter-web/utils"
	"github.com/gorilla/mux"
	"html/template"
)

type Page struct {
	Title   string
	Message string
}

func HomePageHandler(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles(
		"templates/homepage.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/issues.html",
		"templates/contribute.html")

	p := &Page{Title: "Software Craftsmanship Newsletter"}

	if err := t.ExecuteTemplate(res, "homepage", p); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ContributeHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	link := req.Form["link"]
	if err := utils.SendEmail(link[0]); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/thankyou", http.StatusSeeOther)

}

func thankYouHandler(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles(
		"templates/thankyou.html",
		"templates/header.html",
		"templates/footer.html")

	p := &Page{
		Title:   "Software Craftsmanship Newsletter",
		Message: "Thank you very much !"}

	if err := t.ExecuteTemplate(res, "thankyou", p); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePageHandler)
	r.HandleFunc("/submit", ContributeHandler)
	r.HandleFunc("/thankyou", thankYouHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	return r
}

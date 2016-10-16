package routing

import (
	"net/http"

	"html/template"

	"github.com/alebaffa/swcraftnewsletter/mail"
	"github.com/gorilla/mux"
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
	if err := mail.Send(link[0]); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/thankyou", http.StatusSeeOther)

}

func ThankYouHandler(res http.ResponseWriter, req *http.Request) {
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
	r.HandleFunc("/thankyou", ThankYouHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	return r
}

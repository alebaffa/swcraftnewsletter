package routing

import (
	"fmt"
	"net/http"

	"github.com/alebaffa/newsletter-web/utils"
	"github.com/gorilla/mux"
)

func HomePageHandler(res http.ResponseWriter, req *http.Request) {
	utils.ExecuteTemplate("homepage", res)
}

func ContributeHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	link := req.Form["link"]
	fmt.Println("link:", link[0])
	if err := utils.SendEmail(link[0]); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/thankyou", http.StatusSeeOther)

}

func ThankYouHandler(res http.ResponseWriter, req *http.Request) {
	utils.ExecuteTemplate("contribute", res)
}

func NewMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePageHandler)
	r.HandleFunc("/submit", ContributeHandler)
	r.HandleFunc("/thankyou", ThankYouHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	http.Handle("/", r)
	return r
}

package main

import (
	"net/http"

	"github.com/alebaffa/newsletter-web/routing"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routing.HomePageHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

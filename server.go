package main

import (
	"net/http"

	"github.com/alebaffa/newsletter-web/routing"
)

func main() {
	http.ListenAndServe(":8080", routing.NewMux())
}

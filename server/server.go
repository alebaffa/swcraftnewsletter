package main

import (
	"net/http"

	"github.com/alebaffa/newsletter-web/routing"
)

func main() {
	http.ListenAndServe(":3000", routing.NewMux())
}

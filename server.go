package main

import (
	"net/http"

	"github.com/alebaffa/swcraftnewsletter/routing"
)

func main() {
	http.ListenAndServe(":8080", routing.NewMux())
}

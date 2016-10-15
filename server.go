package main

import (
	"net/http"
	"os"

	"github.com/alebaffa/swcraftnewsletter/routing"
)

func main() {
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, routing.NewMux())
}

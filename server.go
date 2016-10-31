package main

import (
	"net/http"
	"os"

	"github.com/alebaffa/swcraftnewsletter/routing"
)

func main() {
	port := os.Getenv("PORT") //heroku
	//port := "8080"
	http.ListenAndServe(":"+port, routing.NewMux())
}

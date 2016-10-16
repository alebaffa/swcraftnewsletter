package main

import (
	"net/http"

	"github.com/alebaffa/swcraftnewsletter/routing"
)

func main() {
	//port := os.Getenv("PORT") //heroku
	port := "8080"
	http.ListenAndServe(":"+port, routing.NewMux())
}

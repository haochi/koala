package main

import (
	"github.com/haochi/koala"
	"net/http"
)

func main() {
	app := koala.New()

	app.AddRouteHandler("GET", "/static/{path}", http.StripPrefix("/static/", http.FileServer(http.Dir("/tmp"))))

	panic(http.ListenAndServe(":8080", app))
}

package main

import (
	"github.com/haochi/koala"
	"net/http"
	"fmt"
)

func main() {
	app := koala.New()

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	app.Get("/~{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user %s reporting for duty!", koala.Param(r, "id"))
	})

	app.ListenAndServe(":8080")
}

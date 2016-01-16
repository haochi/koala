package main

import (
	"fmt"
	"net/http"
	"github.com/haochi/koala"
)

func main() {
	app := koala.New()
	app2 := koala.New()

	app.ServeStaticFiles("/static", "/tmp")
	app.Get("/~{id}", func (writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "user %s reporting for duty!", koala.Param(request, "id"))
	})
	app.Get("/~{id}/hello", func (writer http.ResponseWriter, request *http.Request) {
		go fmt.Fprintf(writer, "hello %s!", koala.Param(request, "id"))
	})

	app2.Get("/", func (writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "serving the second app at a different port: ", request.Host)
	})

	// 	You can start an app like this
	// 	app.ListenAndServe(":8080")

	// 	Or if you want to run multiple apps at different ports
	// 	you can use the helper `koala.ListenAndServe` method
	koala.ListenAndServe(map[*koala.Mux]string {
		app: ":8080",
		app2: ":8081",
	})
}

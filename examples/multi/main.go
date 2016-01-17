package main

import (
	"fmt"
	"net/http"
	"github.com/haochi/koala"
)

func main() {
	app := koala.New()
	app2 := koala.New()

	app.Get("/~{id}", func (writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "user %s reporting for duty!", koala.Param(request, "id"))
	})

	app.Get("/~{id}/hello", func (writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello %s!", koala.Param(request, "id"))
	})

	app2.Get("/", func (writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "serving the second app at a different port: ", request.Host)
	})

	ListenAndServe(map[*koala.Mux]string {
		app: ":8080",
		app2: ":8081",
	})
}

func ListenAndServe(muxes map[*koala.Mux]string) {
	index, lastIndex := 0, len(muxes)-1
	for mux, address := range muxes {
		if index == lastIndex {
			panic(http.ListenAndServe(address, mux))
		} else {
			go func(mux *koala.Mux, address string) {
				panic(http.ListenAndServe(address, mux))
			}(mux, address)
		}
		index++
	}
}

package koala

import (
	"net/http"
	"testing"
)

func TestCreateNewApp(t *testing.T) {
	mux := New()
	if len(mux.routeMap) != 0 {
		t.Error()
	}
}

func TestAddRoute(t *testing.T) {
	mux := New()
	count := 10

	for i := 0; i < count; i++ {
		mux.AddRoute("GET", "/", simpleHandler)
	}

	if len(mux.routeMap["GET"]) != count {
		t.Error()
	}
}

func TestHTTPVerbs(t *testing.T) {
	mux := New()

	mux.Get("/", simpleHandler)
	mux.Post("/", simpleHandler)
	mux.Put("/", simpleHandler)
	mux.Delete("/", simpleHandler)

	if len(mux.routeMap) != 4 {
		t.Error()
	}
}

func TestRouteOrder(t *testing.T) {
	mux := New()

	mux.Get("/{id}", simpleHandler)
	mux.Get("/{id}/hello", simpleHandler)

	_, params := getHandlerAndParamsForRequest(mux, "GET", "/haochi/hello")

	if params["id"] != "haochi" {
		t.Error(params["id"])
	}
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {

}

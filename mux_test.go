package koala

import (
	"net/http"
	"testing"
)

func TestCreateNewApp(t *testing.T) {
	mux := New()
	if len(mux.Routes) != 0 {
		t.Error()
	}

	if mux.ServeMux == nil {
		t.Error()
	}
}

func TestAddRoute(t *testing.T) {
	mux := New()
	count := 10

	for i := 0; i < count; i++ {
		mux.AddRoute("GET", "/", simpleHandler)
	}

	if len(mux.Routes) != count {
		t.Error()
	}
}

func TestHTTPVerbs(t *testing.T) {
	mux := New()

	mux.Get("/", simpleHandler)
	mux.Post("/", simpleHandler)
	mux.Post("/", simpleHandler)
	mux.Delete("/", simpleHandler)

	if len(mux.Routes) != 4 {
		t.Error()
	}
}

func TestServeStaticFiles(t *testing.T) {
	mux := New()
	mux.ServeStaticFiles("/", "")
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {

}

package koala

import "testing"

func TestNewRouteWithParam(t *testing.T) {
	method, path := "GET", "/{id}"
	route := NewRoute(method, path, simpleHandler)

	if route.method != "GET" || route.path != path || route.pathRegExp == nil {
		t.Error()
	}
}

func TestNewRouteWithoutParam(t *testing.T) {
	method, path := "GET", "/id"
	route := NewRoute(method, path, simpleHandler)

	if route.pathRegExp != nil {
		t.Error()
	}
}

func TestResolve(t *testing.T) {
	route := NewRoute("GET", "/{id}", simpleHandler)
	resolved, params := route.resolve("/haochi", "GET")

	if !resolved || len(params["id"]) == 0 {
		t.Error(resolved, params)
	}
}

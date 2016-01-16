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

func TestResolveSimpleTerminatingMatch(t *testing.T) {
	route := NewRoute("GET", "/{id}", simpleHandler)
	resolved, params := route.resolve("/haochi", "GET")

	if !resolved || len(params) == 0 || params["id"] != "haochi" {
		t.Error(resolved, params)
	}
}

func TestResolveNonTerminatingMatch(t *testing.T) {
	route := NewRoute("GET", "/{id}-orange", simpleHandler)
	resolved, params := route.resolve("/haochi-orange", "GET")

	if !resolved || len(params) == 0 || params["id"] != "haochi" {
		t.Error(resolved, params)
	}
}

func TestResolveMultiValueMatch(t *testing.T) {
	route := NewRoute("POST", "/{id}-{name}", simpleHandler)
	resolved, params := route.resolve("/1-haochi", "POST")

	if !resolved || len(params) != 2 || params["id"] != "1" || params["name"] != "haochi" {
		t.Error(resolved, params)
	}
}

func TestResolveNonAsciiMatch(t *testing.T) {
	route := NewRoute("PUT", "/{id}-{name}", simpleHandler)
	resolved, params := route.resolve("/1-世界", "PUT")

	if !resolved || len(params) != 2 || params["id"] != "1" || params["name"] != "世界" {
		t.Error(resolved, params)
	}
}

func TestDoNotResolveInvalidPath(t *testing.T) {
	route := NewRoute("PUT", "/~{id}/asd", simpleHandler)
	resolved, params := route.resolve("/~haochi/abc", "PUT")

	if resolved {
		t.Error(resolved, params)
	}
}

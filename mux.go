package koala

import (
	"net/http"
)

type Mux struct {
	Routes          []*Route
	NotFoundHandler http.Handler
}

const slash = "/"

// Add a new route
func (mux *Mux) AddRoute(method string, path string, handler http.HandlerFunc) *Mux {
	route := NewRoute(method, path, handler)
	mux.Routes = append(mux.Routes, route)
	return mux
}

// Short hand for AddRoute("GET", path, handler)
func (mux *Mux) Get(path string, handler http.HandlerFunc) *Mux {
	return mux.AddRoute("GET", path, handler)
}

// Short hand for AddRoute("POST", path, handler)
func (mux *Mux) Post(path string, handler http.HandlerFunc) *Mux {
	return mux.AddRoute("POST", path, handler)
}

// Short hand for AddRoute("PUT", path, handler)
func (mux *Mux) Put(path string, handler http.HandlerFunc) *Mux {
	return mux.AddRoute("PUT", path, handler)
}

// Short hand for AddRoute("DELETE", path, handler)
func (mux *Mux) Delete(path string, handler http.HandlerFunc) *Mux {
	return mux.AddRoute("DELETE", path, handler)
}

// Serve HTTP
func (mux *Mux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handlerFunc, params := mux.getHandlerAndParamsForRequest(request.Method, request.URL.Path)
	setRequestVars(request, params)
	handlerFunc(writer, request)
	go deleteRequestVars(request)
}

func (mux *Mux) getHandlerAndParamsForRequest(method, path string) (http.HandlerFunc, map[string]string) {
	var resolvedHandler = mux.NotFoundHandler.ServeHTTP
	var resolvedParams map[string]string = nil
	var pathLength = -1

	for _, route := range mux.Routes {
		resolved, params := route.resolve(path, method)

		if resolved && route.pathLength > pathLength {
			resolvedHandler = route.handler
			resolvedParams = params
			pathLength = route.pathLength
		}
	}

	return resolvedHandler, resolvedParams
}

package koala

import (
	"net/http"
)

const slash = "/"

type Mux struct {
	routeMap        map[string][]*Route
	NotFoundHandler http.Handler
}

// Create a new mux
func New() *Mux {
	mux := &Mux{}
	mux.routeMap = make(map[string][]*Route)
	mux.NotFoundHandler = http.NotFoundHandler()
	return mux
}

// Add a new route
func (mux *Mux) AddRoute(method string, path string, handler http.HandlerFunc) *Mux {
	route := NewRoute(method, path, handler)
	mux.routeMap[method] = append(mux.routeMap[method], route)
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
	handlerFunc, params := getHandlerAndParamsForRequest(mux, request.Method, request.URL.Path)
	setRequestVars(request, params)
	handlerFunc(writer, request)
	go deleteRequestVars(request)
}

func getHandlerAndParamsForRequest(mux *Mux, method string, path string) (http.HandlerFunc, map[string]string) {
	var resolvedHandler = mux.NotFoundHandler.ServeHTTP
	var resolvedParams map[string]string = nil
	var pathLength = -1

	for _, route := range mux.routeMap[method] {
		resolved, params := route.resolve(path, method)

		if resolved && route.pathLength > pathLength {
			resolvedHandler = route.handler
			resolvedParams = params
			pathLength = route.pathLength
		}
	}

	return resolvedHandler, resolvedParams
}

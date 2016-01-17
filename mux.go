// koala is a simple multiplexer for go
//
package koala

import (
	"net/http"
)

const slash = "/"

type Mux struct {
	routeMap        map[string][]*Route
	NotFoundHandler http.Handler // Handler when no routes are matched. Defaults to `http.NotFoundHandler`
}

// Create a new mux
func New() *Mux {
	mux := &Mux{}
	mux.routeMap = make(map[string][]*Route)
	mux.NotFoundHandler = http.NotFoundHandler()
	return mux
}

// Add a new route
func (mux *Mux) AddRoute(method string, path string, handlerFunc http.HandlerFunc) *Mux {
	return mux.AddRouteHandler(method, path, http.HandlerFunc(handlerFunc))
}

func (mux *Mux) AddRouteHandler(method string, path string, handler http.Handler) *Mux {
	route := NewRoute(method, path, handler)
	mux.routeMap[method] = append(mux.routeMap[method], route)
	return mux
}

// Serve HTTP
func (mux *Mux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler, params := getHandlerAndParamsForRequest(mux, request.Method, request.URL.Path)
	setRequestVars(request, params)
	handler.ServeHTTP(writer, request)
	go deleteRequestVars(request)
}

func getHandlerAndParamsForRequest(mux *Mux, method string, path string) (http.Handler, map[string]string) {
	var resolvedHandler = mux.NotFoundHandler
	var resolvedParams map[string]string = nil
	var pathLength = -1

	for _, route := range mux.routeMap[method] {
		resolved, staticPath, params := route.resolve(path, method)

		if resolved && route.pathLength > pathLength {
			resolvedHandler = route.handler
			resolvedParams = params
			pathLength = route.pathLength

			if staticPath {
				break
			}
		}
	}

	return resolvedHandler, resolvedParams
}

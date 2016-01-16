package koala

import (
	"net/http"
	"strings"
)

type Mux struct {
	Routes   []*Route
	ServeMux *http.ServeMux
	NotFound http.Handler
}

const slash = "/"

// Add a new route
func (mux *Mux) AddRoute(method string, path string, handler http.HandlerFunc) *Mux {
	route := NewRoute(method, path, handler)
	mux.Routes = append(mux.Routes, route)
	return mux
}

// Serve static files at the given urlPath from local directory directory
func (mux *Mux) ServeStaticFiles(urlPath string, directory string) *Mux {
	path := ensureTrailingSlash(urlPath)
	mux.AddHandler(path, http.StripPrefix(path, http.FileServer(http.Dir(directory))))
	return mux
}

// Add a custom handler
func (mux *Mux) AddHandler(path string, handler http.Handler) *Mux {
	mux.ServeMux.Handle(path, handler)
	return mux
}

// Start running the mux
func (mux *Mux) ListenAndServe(addr string) {
	if len(addr) == 0 {
		panic("`addr` has not been defined.")
	}
	mux.ServeMux.HandleFunc(slash, mux.runner)
	panic(http.ListenAndServe(addr, mux.ServeMux))
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

func ensureTrailingSlash(str string) string {
	if !strings.HasSuffix(str, slash) {
		str += slash
	}
	return str
}

func (mux *Mux) runner(writer http.ResponseWriter, request *http.Request) {
	for _, route := range mux.Routes {
		resolved, params := route.resolve(request.URL.Path, request.Method)
		if resolved {
			setRequestVars(request, params)
			route.handler(writer, request)
			deleteRequestVars(request)
			return
		}
	}

	http.NotFound(writer, request)
}

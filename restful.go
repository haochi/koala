package koala

import "net/http"

// Short hand for AddRoute("GET", path, handlerFunc)
func (mux *Mux) Get(path string, handlerFunc http.HandlerFunc) *Mux {
	return mux.AddRoute("GET", path, handlerFunc)
}

// Short hand for AddRoute("POST", path, handlerFunc)
func (mux *Mux) Post(path string, handlerFunc http.HandlerFunc) *Mux {
	return mux.AddRoute("POST", path, handlerFunc)
}

// Short hand for AddRoute("PUT", path, handlerFunc)
func (mux *Mux) Put(path string, handlerFunc http.HandlerFunc) *Mux {
	return mux.AddRoute("PUT", path, handlerFunc)
}

// Short hand for AddRoute("DELETE", path, handlerFunc)
func (mux *Mux) Delete(path string, handlerFunc http.HandlerFunc) *Mux {
	return mux.AddRoute("DELETE", path, handlerFunc)
}

// Short hand for AddRouteHandler("GET", path, handler)
func (mux *Mux) GetHandler(path string, handler http.Handler) *Mux {
	return mux.AddRouteHandler("GET", path, handler)
}

// Short hand for AddRouteHandler("POST", path, handler)
func (mux *Mux) PostHandler(path string, handler http.Handler) *Mux {
	return mux.AddRouteHandler("POST", path, handler)
}

// Short hand for AddRouteHandler("PUT", path, handler)
func (mux *Mux) PutHandler(path string, handler http.Handler) *Mux {
	return mux.AddRouteHandler("PUT", path, handler)
}

// Short hand for AddRouteHandler("DELETE", path, handler)
func (mux *Mux) DeleteHandler(path string, handler http.Handler) *Mux {
	return mux.AddRouteHandler("DELETE", path, handler)
}

// koala is a lightweight multiplexer for Go
//
package koala

import "net/http"

// Create a new mux
func New() *Mux {
	mux := &Mux{}
	mux.NotFoundHandler = http.NotFoundHandler()
	return mux
}

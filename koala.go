// koala is a lightweight multiplexer for Go
//
package koala

import "net/http"

// Create a new mux
func New() *Mux {
	mux := &Mux{}
	mux.ServeMux = http.NewServeMux()
	return mux
}

func ListenAndServe(muxes map[*Mux]string) {
	index, lastIndex := 0, len(muxes)-1
	for mux, address := range muxes {
		if index == lastIndex {
			mux.ListenAndServe(address)
		} else {
			go func(mux *Mux, port string) {
				mux.ListenAndServe(port)
			}(mux, address)
		}
		index++
	}
}

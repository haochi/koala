// koala is a lightweight multiplexer for Go
//
package koala

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

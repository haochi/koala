package koala

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetRequestVars(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key, val := "hello", "world"
		params := make(map[string]string)
		params[key] = val
		setRequestVars(r, params)

		if len(GetPathParams(r)) != 1 || GetPathParam(r, key) != val {
			t.Error()
		}

		deleteRequestVars(r)

		if len(GetPathParams(r)) != 0 {
			t.Error()
		}
	}))
	defer ts.Close()
	http.Get(ts.URL)
}

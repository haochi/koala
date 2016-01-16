package koala

import (
	"net/http"
	"sync"
)

var locker = sync.RWMutex{}
var vars = make(map[*http.Request]map[string]string)

// Get the resolved route parameter with the given key
func Param(request *http.Request, key string) string {
	pathVars := Params(request)
	return pathVars[key]
}

// Get all the resolved route parameters
func Params(request *http.Request) map[string]string {
	locker.RLock()
	params := vars[request]
	locker.RUnlock()
	return params
}

func setRequestVars(request *http.Request, params map[string]string) {
	locker.Lock()
	vars[request] = params
	locker.Unlock()
}

func deleteRequestVars(request *http.Request) {
	locker.Lock()
	delete(vars, request)
	locker.Unlock()
}

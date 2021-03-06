package koala

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

var pathToRegExpReplacer = `(?P<%s>.*)`
var paramRegExp = regexp.MustCompile(fmt.Sprintf(`\{(%s)\}`, `[^\}]+`))

type Route struct {
	path       string
	pathRegExp *regexp.Regexp
	pathLength int
	method     string
	handler    http.Handler
}

func NewRoute(method string, path string, handler http.Handler) *Route {
	route := &Route{path: path, method: method, handler: handler}
	initRoutePathRegExp(route)
	route.pathLength = len(path)
	return route
}

func (route *Route) resolve(path string, method string) (resolved bool, staticPath bool, params map[string]string) {
	if route.method == method {
		if route.path == path {
			resolved = true
			staticPath = true
		} else if route.pathRegExp != nil {
			resolved, params = resolvePathParams(path, route.pathRegExp)
		}
	}
	return
}

func initRoutePathRegExp(route *Route) {
	pathRegExp := paramRegExp.ReplaceAllStringFunc(route.path, pathToRegexReplacer)
	if route.path != pathRegExp {
		regExp, err := regexp.Compile(absoluteWrapPathRegExp(pathRegExp))
		if err != nil {
			panic(errors.New(fmt.Sprintln("Invalid path definition: %s", route.path)))
		}
		route.pathRegExp = regExp
	}
}

func absoluteWrapPathRegExp(pathRegExp string) string {
	return fmt.Sprintf(`^%s/?$`, pathRegExp)
}

func pathToRegexReplacer(param string) string {
	return fmt.Sprintf(pathToRegExpReplacer, param[1:len(param)-1])
}

func resolvePathParams(path string, pathRegExp *regexp.Regexp) (bool, map[string]string) {
	match := pathRegExp.FindStringSubmatch(path)
	if len(match) > 0 {
		params := make(map[string]string)
		for i, name := range pathRegExp.SubexpNames() {
			if i != 0 {
				params[name] = match[i]
			}
		}

		return true, params
	}
	return false, nil
}

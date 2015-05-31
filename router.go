package main

import (
	"log"
	"net/http"
	"regexp"
)

type RouteParams map[string]string

// http://blog.golang.org/error-handling-and-go
type Handler func(http.ResponseWriter, *http.Request, RouteParams) error

// -- Route
type Route struct {
	method     string
	pattern    *regexp.Regexp
	paramNames []string
	handler    Handler
}

func NewRoute(method string, patternString string, handler Handler) *Route {
	pattern, paramNames := parsePattern(patternString)
	return &Route{method, pattern, paramNames, handler}
}

func parsePattern(patternString string) (*regexp.Regexp, []string) {
	// Extract param names.
	// "/foo/{fooId}/bar/{barId}" -> []string{"fooId", "barId"}
	paramPattern := regexp.MustCompile("\\{([^\\}]+)\\}")
	matches := paramPattern.FindAllStringSubmatch(patternString, -1)
	paramNames := make([]string, len(matches))
	for i, value := range matches {
		paramNames[i] = value[1]
	}
	// Create a regexp to extract params from path.
	// "/foo/{fooId}/bar/{barId}" -> "/foo/([a-zA-Z0-9]+)/bar/([a-zA-Z0-9]+)"
	pathPatternString := paramPattern.ReplaceAllString(patternString, "([a-zA-Z0-9]+)")
	pattern := regexp.MustCompile("^" + pathPatternString + "$")

	return pattern, paramNames
}

// Check if the route matches given method and path
func (route *Route) matches(method string, path string) (bool, RouteParams) {
	matched := route.method == method && route.pattern.MatchString(path)
	if !matched {
		return false, nil
	}

	parts := route.pattern.FindStringSubmatch(path)
	params := make(RouteParams)
	for i, paramName := range route.paramNames {
		params[paramName] = parts[i+1]
	}

	return matched, params
}

// -- Router
type Router struct {
	routes []*Route
}

func NewRouter() *Router {
	return &Router{}
}

// Router implements http.ServeMux.
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)
	route, params := router.findRoute(r.Method, r.URL.Path)
	if route == nil {
		http.NotFound(w, r)
		return
	}

	err := route.handler(w, r, params)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (router *Router) findRoute(method string, pattern string) (*Route, RouteParams) {
	for _, route := range router.routes {
		matched, params := route.matches(method, pattern)
		if matched {
			return route, params
		}
	}
	return nil, nil
}

// Add a handler for specified method.
func (router *Router) AddHandler(method string, pattern string, handler Handler) {
	route := NewRoute(method, pattern, handler)
	router.routes = append(router.routes, route)
}

// Add a handler for GET method.
func (router *Router) GET(pattern string, handler Handler) {
	router.AddHandler("GET", pattern, handler)
}

// Add a hanlder for POST method.
func (router *Router) POST(pattern string, handler Handler) {
	router.AddHandler("POST", pattern, handler)
}

// Add a handler for PUT method.
func (router *Router) PUT(pattern string, handler Handler) {
	router.AddHandler("PUT", pattern, handler)
}

// Add a handler for DELETE method.
func (router *Router) DELETE(pattern string, handler Handler) {
	router.AddHandler("DELETE", pattern, handler)
}

// Add a handler for HEAD method.
func (router *Router) HEAD(pattern string, handler Handler) {
	router.AddHandler("HEAD", pattern, handler)
}

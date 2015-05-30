package main

import (
	"log"
	"net/http"
)

// http://blog.golang.org/error-handling-and-go
type Handler func(http.ResponseWriter, *http.Request) error

// -- Route
type Route struct {
	method  string
	pattern string
	handler Handler
}

func NewRoute(method string, pattern string, handler Handler) *Route {
	return &Route{method, pattern, handler}
}

func (route *Route) matches(method string, pattern string) bool {
	return route.method == method && route.pattern == pattern
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
	route := router.findRoute(r.Method, r.URL.Path)
	if route == nil {
		http.NotFound(w, r)
		return
	}

	err := route.handler(w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (router *Router) findRoute(method string, pattern string) *Route {
	for _, route := range router.routes {
		if route.matches(method, pattern) {
			return route
		}
	}
	return nil
}

// Add a handler for specified method.
func (router *Router) AddHandler(method string, pattern string, handler Handler) {
	route := &Route{method, pattern, handler}
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

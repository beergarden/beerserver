package main

import (
	"net/http"
	"testing"
)

// -- Tests
func Test_Route_matches(t *testing.T) {
	route := NewRoute("GET", "/foo", testHandler)
	expectToMatch(t, route, "GET", "/foo")
	expectNotToMatch(t, route, "GET", "/fo")
	expectNotToMatch(t, route, "GET", "/foos")
}

func Test_Route_matches_root(t *testing.T) {
	route := NewRoute("GET", "/", testHandler)
	expectToMatch(t, route, "GET", "/")
	expectNotToMatch(t, route, "GET", "/foo")
}

// -- Utils
func expectToMatch(t *testing.T, route *Route, method string, pattern string) {
	if !route.matches(method, pattern) {
		t.Errorf("Expected to match %v %v but didn't", method, pattern)
	}
}

func expectNotToMatch(t *testing.T, route *Route, method string, pattern string) {
	if route.matches(method, pattern) {
		t.Errorf("Expected not to match %v %v but did", method, pattern)
	}
}

func testHandler(http.ResponseWriter, *http.Request) error {
	return nil
}

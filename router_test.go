package main

import (
	"net/http"
	"reflect"
	"testing"
)

// -- Tests
func Test_Route_matches(t *testing.T) {
	route := NewRoute("GET", "/foo", testHandler)
	expectToMatch(t, route, "GET", "/foo", RouteParams{})
	expectNotToMatch(t, route, "GET", "/fo")
	expectNotToMatch(t, route, "POST", "/foo")
	expectNotToMatch(t, route, "GET", "/foos")
}

func Test_Route_matches_root(t *testing.T) {
	route := NewRoute("GET", "/", testHandler)
	expectToMatch(t, route, "GET", "/", RouteParams{})
	expectNotToMatch(t, route, "GET", "/foo")
}

func Test_Route_matches_params(t *testing.T) {
	route := NewRoute("GET", "/foo/{fooId}/bar/{barId}", testHandler)
	expectToMatch(t, route, "GET", "/foo/123/bar/234", RouteParams{"fooId": "123", "barId": "234"})
	expectNotToMatch(t, route, "GET", "/foo/123/bar")
}

// -- Utils
func expectToMatch(t *testing.T, route *Route, method string, pattern string, params RouteParams) {
	matched, actualParams := route.matches(method, pattern)
	if !matched {
		t.Errorf("Expected to match %v %v but didn't", method, pattern)
	}
	if !reflect.DeepEqual(params, actualParams) {
		t.Errorf("Expected params %v but got %v", params, actualParams)
	}
}

func expectNotToMatch(t *testing.T, route *Route, method string, pattern string) {
	matched, _ := route.matches(method, pattern)
	if matched {
		t.Errorf("Expected not to match %v %v but did", method, pattern)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request, params RouteParams) error {
	return nil
}

package goapi

import "net/http"

type Router struct {
	*Group
}

// NewRouter creates and returns a new instance of Router.
// The Router is responsible for handling incoming HTTP requests and routing them to the appropriate handlers.
//
// Return:
// - *Router: A pointer to the newly created Router instance.
func NewRouter() *Router {
	return &Router{
		Group: &Group{
			prefix:     "",
			middleware: make([]MiddlewareFunc, 0),
			routes:     make([]route, 0),
			subgroups:  make([]*Group, 0),
			parent:     nil,
		},
	}
}

// ServerHTTP is the HTTP handler function for the Router. It matches incoming requests
// to the registered routes and executes the corresponding handlers. If no matching route is found,
// it returns a 404 Not Found response.
//
// Parameters:
// - w: http.ResponseWriter to write the response.
// - req: *http.Request representing the incoming HTTP request.
//
// Return:
// - None.
func (r *Router) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	if !r.handleRequest(w, req) {
		http.NotFound(w, req)
	}
}

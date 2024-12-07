package goapi

import "net/http"

type Router struct {
	*Group
}

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

func (r *Router) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	if !r.handleRequest(w, req) {
		http.NotFound(w, req)
	}
}

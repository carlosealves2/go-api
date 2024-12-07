package goapi

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

var paramsKey = contextKey("route_params")

type Group struct {
	prefix     string
	middleware []MiddlewareFunc
	routes     []route
	subgroups  []*Group
	parent     *Group
}

func (g *Group) Group(prefix string) *Group {
	subgroup := &Group{
		prefix:     g.prefix + prefix,
		middleware: make([]MiddlewareFunc, 0),
		routes:     make([]route, 0),
		subgroups:  make([]*Group, 0),
		parent:     g,
	}
	g.subgroups = append(g.subgroups, subgroup)
	return subgroup
}

func (g *Group) Use(middleware ...MiddlewareFunc) {
	g.middleware = append(g.middleware, middleware...)
}

func (g *Group) Handle(method, pattern string, handler HandlerFunc) {
	fullPattern := strings.TrimRight(g.prefix, "/") + "/" + strings.TrimLeft(pattern, "/")
	regexPattern, paramNames := parsePattern(fullPattern)
	g.routes = append(g.routes, route{
		method:     method,
		pattern:    regexPattern,
		paramNames: paramNames,
		handler:    handler,
	})
}

func (g *Group) GET(pattern string, handler HandlerFunc) {
	g.Handle("GET", pattern, handler)
}

func (g *Group) POST(pattern string, handler HandlerFunc) {
	g.Handle("POST", pattern, handler)
}

func (g *Group) PUT(pattern string, handler HandlerFunc) {
	g.Handle("PUT", pattern, handler)
}

func (g *Group) DELETE(pattern string, handler HandlerFunc) {
	g.Handle("DELETE", pattern, handler)
}

func (g *Group) PATCH(pattern string, handler HandlerFunc) {
	g.Handle("PATCH", pattern, handler)
}

func (g *Group) HEAD(pattern string, handler HandlerFunc) {
	g.Handle("HEAD", pattern, handler)
}

func (g *Group) OPTIONS(pattern string, handler HandlerFunc) {
	g.Handle("OPTIONS", pattern, handler)
}

func (g *Group) handleRequest(w http.ResponseWriter, r *http.Request) bool {
	for _, requestedRoute := range g.routes {
		if r.Method != requestedRoute.method {
			continue
		}

		matches := requestedRoute.pattern.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			params := make(map[string]string)
			for i, name := range requestedRoute.paramNames {
				params[name] = matches[i+1]
			}

			ctx := context.WithValue(r.Context(), paramsKey, params)
			r = r.WithContext(ctx)

			middlewares := g.collectMiddlewares()

			finalHandler := requestedRoute.handler
			for i := len(middlewares) - 1; i >= 0; i-- {
				finalHandler = middlewares[i](finalHandler)
			}

			finalHandler(w, r)
			return true
		}
	}

	for _, subgroup := range g.subgroups {
		if subgroup.handleRequest(w, r) {
			return true
		}
	}
	w.WriteHeader(http.StatusNotFound)
	return false
}

func (g *Group) collectMiddlewares() []MiddlewareFunc {
	var middlewares []MiddlewareFunc
	if g.parent != nil {
		middlewares = g.parent.collectMiddlewares()
	}

	middlewares = append(middlewares, g.middleware...)
	return middlewares
}

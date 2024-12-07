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

// Group creates a new subgroup with the specified prefix and adds it to the current group.
// The prefix will be appended to the parent group's prefix.
//
// prefix: The prefix to be added to the parent group's prefix.
//
// Returns: A pointer to the newly created subgroup.
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

// Use adds one or more middleware functions to the current group.
// The middleware functions will be executed in the order they are added.
//
// middleware: A variadic parameter that accepts one or more middleware functions.
// These middleware functions will be appended to the current group's middleware list.
//
// Example:
//
//	api := goapi.New()
//	api.Use(func(next goapi.HandlerFunc) goapi.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			fmt.Println("Before request")
//			next(w, r)
//			fmt.Println("After request")
//		}
//	})
func (g *Group) Use(middleware ...MiddlewareFunc) {
	g.middleware = append(g.middleware, middleware...)
}

// Handle adds a new route to the current group with the specified method, pattern, and handler function.
// The full pattern will be constructed by appending the pattern to the parent group's prefix.
// The pattern will be parsed to extract any dynamic parameters, which will be available in the request context.
//
// method: The HTTP method for the route (e.g., "GET", "POST", "PUT", "DELETE", etc.).
// pattern: The pattern for the route, which may contain dynamic parameters enclosed in curly braces (e.g., "/users/{id}").
// handler: The handler function to be executed when the route is matched.
//
// Example:
//
//	api := goapi.New()
//	usersGroup := api.Group("/users")
//	usersGroup.Handle("GET", "/{id}", func(w http.ResponseWriter, r *http.Request) {
//		params := r.Context().Value(goapi.paramsKey).(map[string]string)
//		userID := params["id"]
//		// ...
//	})
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

// GET is a shortcut method for adding a new route with the HTTP method "GET" to the current group.
// The full pattern will be constructed by appending the pattern to the parent group's prefix.
// The pattern will be parsed to extract any dynamic parameters, which will be available in the request context.
//
// pattern: The pattern for the route, which may contain dynamic parameters enclosed in curly braces (e.g., "/users/{id}").
// handler: The handler function to be executed when the route is matched.
//
// Example:
//
//	api := goapi.New()
//	usersGroup := api.Group("/users")
//	usersGroup.GET("/{id}", func(w http.ResponseWriter, r *http.Request) {
//		params := r.Context().Value(goapi.paramsKey).(map[string]string)
//		userID := params["id"]
//		// ...
//	})
func (g *Group) GET(pattern string, handler HandlerFunc) {
	g.Handle("GET", pattern, handler)
}

// POST adds a new route with the HTTP method "POST" to the current group.
// The full pattern will be constructed by appending the pattern to the parent group's prefix.
// The pattern will be parsed to extract any dynamic parameters, which will be available in the request context.
//
// Parameters:
// - pattern: The pattern for the route, which may contain dynamic parameters enclosed in curly braces (e.g., "/users/{id}").
// - handler: The handler function to be executed when the route is matched.
//
// Returns:
// This function does not return any value. It adds a new route to the current group.
func (g *Group) POST(pattern string, handler HandlerFunc) {
	g.Handle("POST", pattern, handler)
}

// PUT adds a new route with the HTTP method "PUT" to the current group.
// The full pattern will be constructed by appending the pattern to the parent group's prefix.
// The pattern will be parsed to extract any dynamic parameters, which will be available in the request context.
//
// Parameters:
// - pattern: The pattern for the route, which may contain dynamic parameters enclosed in curly braces (e.g., "/users/{id}").
// - handler: The handler function to be executed when the route is matched.
//
// Returns:
// This function does not return any value. It adds a new route to the current group.
func (g *Group) PUT(pattern string, handler HandlerFunc) {
	g.Handle("PUT", pattern, handler)
}

// DELETE adds a new route with the HTTP method "DELETE" to the current group.
// The full pattern will be constructed by appending the pattern to the parent group's prefix.
// The pattern will be parsed to extract any dynamic parameters, which will be available in the request context.
//
// Parameters:
// - pattern: The pattern for the route, which may contain dynamic parameters enclosed in curly braces (e.g., "/users/{id}").
// - handler: The handler function to be executed when the route is matched.
//
// Returns:
// This function does not return any value. It adds a new route to the current group.
func (g *Group) DELETE(pattern string, handler HandlerFunc) {
	g.Handle("DELETE", pattern, handler)
}

// PATCH adds a new route with the HTTP method "PATCH" to the current group.
// The full pattern will be constructed by appending the pattern to the parent group's prefix.
// The pattern will be parsed to extract any dynamic parameters, which will be available in the request context.
//
// Parameters:
// - pattern: The pattern for the route, which may contain dynamic parameters enclosed in curly braces (e.g., "/users/{id}").
// - handler: The handler function to be executed when the route is matched.
//
// Returns:
// This function does not return any value. It adds a new route to the current group.
func (g *Group) PATCH(pattern string, handler HandlerFunc) {
	g.Handle("PATCH", pattern, handler)
}

// HEAD adds a new route with the HTTP method "HEAD" to the current group.
// The full pattern will be constructed by appending the pattern to the parent group's prefix.
// The pattern will be parsed to extract any dynamic parameters, which will be available in the request context.
//
// Parameters:
// - pattern: The pattern for the route, which may contain dynamic parameters enclosed in curly braces (e.g., "/users/{id}").
// - handler: The handler function to be executed when the route is matched.
//
// Returns:
// This function does not return any value. It adds a new route to the current group.
func (g *Group) HEAD(pattern string, handler HandlerFunc) {
	g.Handle("HEAD", pattern, handler)
}

// OPTIONS adds a new route with the HTTP method "OPTIONS" to the current group.
// The full pattern will be constructed by appending the pattern to the parent group's prefix.
// The pattern will be parsed to extract any dynamic parameters, which will be available in the request context.
//
// Parameters:
// - pattern: The pattern for the route, which may contain dynamic parameters enclosed in curly braces (e.g., "/users/{id}").
// - handler: The handler function to be executed when the route is matched.
//
// Returns:
// This function does not return any value. It adds a new route to the current group.
func (g *Group) OPTIONS(pattern string, handler HandlerFunc) {
	g.Handle("OPTIONS", pattern, handler)
}

// handleRequest processes incoming HTTP requests and matches them to the appropriate route within the group.
// It iterates through the routes and checks if the request method and URL path match. If a match is found,
// it extracts any dynamic parameters from the URL path and adds them to the request context. It then applies
// the middleware functions associated with the group and its parent groups, and finally executes the handler
// function for the matched route. If no match is found, it returns a 404 Not Found response.
//
// Parameters:
// - w: The http.ResponseWriter to write the response.
// - r: The *http.Request containing the incoming HTTP request.
//
// Returns:
//   - A boolean indicating whether a route was matched and processed. If true, the request was handled; otherwise,
//     a 404 Not Found response was sent
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

// collectMiddlewares recursively collects all middleware functions associated with the current group
// and its parent groups. The middleware functions are returned in the order they were added.
//
// Parameters:
// - None
//
// Returns:
// - []MiddlewareFunc: A slice of middleware functions associated with the current group and its parent groups.
func (g *Group) collectMiddlewares() []MiddlewareFunc {
	var middlewares []MiddlewareFunc
	if g.parent != nil {
		middlewares = g.parent.collectMiddlewares()
	}

	middlewares = append(middlewares, g.middleware...)
	return middlewares
}

# Go Router Package

A flexible, tree-structured HTTP router for Go, inspired by the simplicity and flexibility of Express.js. This router supports route groups (subgroups), middlewares at multiple nesting levels, parameterized routes, and regex-based path matching.

**Table of Contents**
- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Route Groups and Subgroups](#route-groups-and-subgroups)
- [Middlewares](#middlewares)
- [Parameterized Routes](#parameterized-routes)
- [Advanced Usage](#advanced-usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Simple API:** Define routes with HTTP methods like `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD` and `OPTIONS`.
- **Parameterized Routes:** Easily extract URL parameters from paths like `/users/:id`.
- **Middlewares:** Attach middlewares at any level—global, group, or subgroup. They are processed in a chain, ensuring modular and reusable logic.
- **Route Groups and Subgroups:** Organize routes in hierarchical groups, each with its own prefix and middlewares.
- **Regex-Based Path Matching:** Routes are matched using compiled regular expressions, allowing for complex URL patterns.
- **HTTP Handler Compatible:** Implemented as an `http.Handler`, so it plugs into the Go `net/http` ecosystem seamlessly.

---

## Installation

Make sure you have Go installed (Go 1.18+ recommended).

```bash
go get github.com/carlosealves2/go-api
```

Then, import it in your code:

```go
import goapi "github.com/carlosealves2/go-api"
```

---

## Quick Start

Create a new router, register routes, and start serving HTTP:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
	
    goapi "github.com/carlosealves2/go-api"
    "github.com/carlosealves2/go-api/middlewares"
)

func main() {
    r := goapi.NewRouter()

    // Global middleware
    r.Use(middlewares.LoggingMiddleware)

    // Root route
    r.GET("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintln(w, "Welcome to the home page!")
    })

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

```

---

## Route Groups and Subgroups

Organize your routes by grouping them under specific prefixes. Each group can have its own middlewares and subgroups.

```go
apiGroup := r.Group("/api")
apiGroup.GET("/status", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "API is up and running!")
})

// Nested subgroup for admin routes
adminGroup := apiGroup.Group("/admin")
adminGroup.GET("/dashboard", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Welcome to the admin dashboard")
})
```

**Key points:**
- Groups inherit the prefix of their parent. For example, `/api/admin/dashboard`.
- Middlewares applied to the parent group also affect subgroups unless overridden or supplemented.

---

## Middlewares

Middlewares are functions that wrap your final handler, allowing you to add logic like logging, authentication, rate limiting, etc.

```go
apiGroup.Use(authMiddleware)

func authMiddleware(next router.HandlerFunc) router.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        token := req.Header.Get("Authorization")
        if token != "valid-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next(w, req, params)
    }
}
```

Middlewares are collected up the group chain. Global middlewares apply to all routes, group-level middlewares apply to all routes within that group and its subgroups.

---

## Parameterized Routes

Extract parameters from the URL path by prefixing them with `:`. These parameters are then accessible from the `params` map in your handler.

```go
apiGroup.GET("/users/:id", func(w http.ResponseWriter, req *http.Request) {
    userID := goapi.ParamsFromContext(r)["id"]
    fmt.Fprintf(w, "User details for ID %s\n", userID)
})
```

**Example:**
- A request to `/api/users/123` sets `params["id"] = "123"`.

---

## Advanced Usage

- **Regex-based Matching:** The router automatically converts patterns like `:id` into `([^/]+)`. You can further customize patterns by adjusting how `parsePattern` handles segments.
- **Complex Group Hierarchies:** Create as many nested groups as your application needs, each adding its own prefix and middlewares.
- **Integration with Standard Library:** The router implements `http.Handler`, so it can be used directly with `http.ListenAndServe` or integrated with other HTTP frameworks and middleware stacks.

---

## Examples

This repository includes an `examples/` directory with a `main.go` that demonstrates:
- Global and group-level middlewares
- Nested groups
- Parameterized routes
- Authentication checks
- Logging

To run the example:

```bash
cd examples
go run main.go
```

Then open `http://localhost:8080/` in your browser.

---

## Contributing

Contributions are welcome! Feel free to open issues, suggest improvements, or submit pull requests. When contributing, please:
- Write clear commit messages
- Add tests or examples for new functionality
- Follow Go best practices and formatting conventions

---

## License

This project is licensed under the [MIT License](LICENSE). You are free to use, modify, and distribute this code as long as you include the original license.


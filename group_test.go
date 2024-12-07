package goapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock Middleware
func mockMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Middleware", "true")
		next(w, r)
	}
}

// Mock Handler
func mockHandler(message string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	}
}

func TestGroupHandleRequest(t *testing.T) {
	root := &Group{}

	// Define routes
	root.GET("/hello", mockHandler("Hello, World!"))
	root.GET("/user/:id", func(w http.ResponseWriter, r *http.Request) {
		params := r.Context().Value(paramsKey).(map[string]string)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User ID: " + params["id"]))
	})

	// Add middleware
	root.Use(mockMiddleware)

	// Test /hello route
	t.Run("GET /hello", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/hello", nil)
		resp := httptest.NewRecorder()

		if !root.handleRequest(resp, req) {
			t.Fatalf("Expected /hello to be handled")
		}

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		if resp.Body.String() != "Hello, World!" {
			t.Errorf("Expected body 'Hello, World!', got '%s'", resp.Body.String())
		}

		if resp.Header().Get("X-Middleware") != "true" {
			t.Errorf("Expected middleware header to be set")
		}
	})

	// Test /user/:id route
	t.Run("GET /user/:id", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/user/42", nil)
		resp := httptest.NewRecorder()

		if !root.handleRequest(resp, req) {
			t.Fatalf("Expected /user/:id to be handled")
		}

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		expectedBody := "User ID: 42"
		if resp.Body.String() != expectedBody {
			t.Errorf("Expected body '%s', got '%s'", expectedBody, resp.Body.String())
		}
	})
}

func TestGroupSubgroups(t *testing.T) {
	root := &Group{}
	api := root.Group("/api")
	api.GET("/status", mockHandler("OK"))

	t.Run("GET /api/status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/status", nil)
		resp := httptest.NewRecorder()

		if !root.handleRequest(resp, req) {
			t.Fatalf("Expected /api/status to be handled")
		}

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		if resp.Body.String() != "OK" {
			t.Errorf("Expected body 'OK', got '%s'", resp.Body.String())
		}
	})
}

func TestGroupMiddleware(t *testing.T) {
	root := &Group{}
	api := root.Group("/api")
	api.Use(mockMiddleware)
	api.GET("/test", mockHandler("Middleware Test"))

	t.Run("GET /api/test with middleware", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/test", nil)
		resp := httptest.NewRecorder()

		if !root.handleRequest(resp, req) {
			t.Fatalf("Expected /api/test to be handled")
		}

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		if resp.Body.String() != "Middleware Test" {
			t.Errorf("Expected body 'Middleware Test', got '%s'", resp.Body.String())
		}

		if resp.Header().Get("X-Middleware") != "true" {
			t.Errorf("Expected middleware header to be set")
		}
	})
}

func TestGroupPOST(t *testing.T) {
	root := &Group{}
	root.POST("/create", mockHandler("Created"))

	t.Run("POST /create", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/create", nil)
		resp := httptest.NewRecorder()

		if !root.handleRequest(resp, req) {
			t.Fatalf("Expected /create to be handled")
		}

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		if resp.Body.String() != "Created" {
			t.Errorf("Expected body 'Created', got '%s'", resp.Body.String())
		}
	})
}

func TestGroupPUT(t *testing.T) {
	root := &Group{}
	root.PUT("/update", mockHandler("Updated"))

	t.Run("PUT /update", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/update", nil)
		resp := httptest.NewRecorder()

		if !root.handleRequest(resp, req) {
			t.Fatalf("Expected /update to be handled")
		}

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		if resp.Body.String() != "Updated" {
			t.Errorf("Expected body 'Updated', got '%s'", resp.Body.String())
		}
	})
}

func TestGroupDELETE(t *testing.T) {
	root := &Group{}
	root.DELETE("/delete", mockHandler("Deleted"))

	t.Run("DELETE /delete", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/delete", nil)
		resp := httptest.NewRecorder()

		if !root.handleRequest(resp, req) {
			t.Fatalf("Expected /delete to be handled")
		}

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		if resp.Body.String() != "Deleted" {
			t.Errorf("Expected body 'Deleted', got '%s'", resp.Body.String())
		}
	})
}

func TestGroupPATCH(t *testing.T) {
	root := &Group{}
	root.PATCH("/modify", mockHandler("Modified"))

	t.Run("PATCH /modify", func(t *testing.T) {
		req := httptest.NewRequest("PATCH", "/modify", nil)
		resp := httptest.NewRecorder()

		if !root.handleRequest(resp, req) {
			t.Fatalf("Expected /modify to be handled")
		}

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		if resp.Body.String() != "Modified" {
			t.Errorf("Expected body 'Modified', got '%s'", resp.Body.String())
		}
	})
}

func TestGroupHEAD(t *testing.T) {
	root := &Group{}
	root.HEAD("/head", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	t.Run("HEAD /head", func(t *testing.T) {
		req := httptest.NewRequest("HEAD", "/head", nil)
		resp := httptest.NewRecorder()

		if !root.handleRequest(resp, req) {
			t.Fatalf("Expected /head to be handled")
		}

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		if resp.Body.String() != "" {
			t.Errorf("Expected empty body for HEAD request, got '%s'", resp.Body.String())
		}
	})
}

func TestGroupOPTIONS(t *testing.T) {
	root := &Group{}
	root.OPTIONS("/options", mockHandler("Options OK"))

	t.Run("OPTIONS /options", func(t *testing.T) {
		req := httptest.NewRequest("OPTIONS", "/options", nil)
		resp := httptest.NewRecorder()

		if !root.handleRequest(resp, req) {
			t.Fatalf("Expected /options to be handled")
		}

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		if resp.Body.String() != "Options OK" {
			t.Errorf("Expected body 'Options OK', got '%s'", resp.Body.String())
		}
	})
}

func TestGroupHandleRequestMethodMismatch(t *testing.T) {
	root := &Group{}
	root.GET("/hello", mockHandler("Hello, World!"))

	t.Run("Method Mismatch - POST /hello", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/hello", nil) // Incorrect method
		resp := httptest.NewRecorder()

		handled := root.handleRequest(resp, req)

		if handled {
			t.Fatalf("Expected handleRequest to return false for method mismatch")
		}

		if resp.Code != http.StatusNotFound {
			t.Errorf("Expected status code to remain 404, got %d", resp.Code)
		}
	})
}

func TestGroupHandleRequestNoMatch(t *testing.T) {
	root := &Group{}
	root.GET("/hello", mockHandler("Hello, World!"))

	t.Run("No Match - GET /nonexistent", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/nonexistent", nil) // Non-existent route
		resp := httptest.NewRecorder()

		handled := root.handleRequest(resp, req)

		if handled {
			t.Fatalf("Expected handleRequest to return false for non-existent route")
		}

		if resp.Code != http.StatusNotFound {
			t.Errorf("Expected status code to remain 404, got %d", resp.Code)
		}
	})
}

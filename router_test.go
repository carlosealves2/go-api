package goapi

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func TestRouterServerHTTP(t *testing.T) {
	router := NewRouter()

	// Adiciona uma rota de teste
	router.routes = append(router.routes, route{
		method:  http.MethodGet,
		pattern: regexp.MustCompile("^/test$"),
		handler: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, Test!"))
		},
	})

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Route",
			method:         http.MethodGet,
			path:           "/test",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, Test!",
		},
		{
			name:           "Invalid Route",
			method:         http.MethodGet,
			path:           "/invalid",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Cria a requisição
			req := httptest.NewRequest(test.method, test.path, nil)
			// Cria o recorder para capturar a resposta
			rec := httptest.NewRecorder()

			// Chama o handler
			router.ServerHTTP(rec, req)

			// Verifica o status da resposta
			if rec.Code != test.expectedStatus {
				t.Errorf("expected status %d, got %d", test.expectedStatus, rec.Code)
			}

			// Verifica o corpo da resposta
			if strings.TrimSpace(rec.Body.String()) != strings.TrimSpace(test.expectedBody) {
				t.Errorf("expected body '%s', got '%s'", test.expectedBody, rec.Body.String())
			}
		})
	}
}

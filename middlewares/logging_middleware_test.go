package middlewares

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}

	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)
	defer log.SetOutput(nil)

	handler := LoggingMiddleware(mockHandler)

	req := httptest.NewRequest("GET", "/test-path", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	expectedBody := "OK"
	if strings.TrimSpace(rec.Body.String()) != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rec.Body.String())
	}

	expectedLog := "GET /test-path"
	if !strings.Contains(logOutput.String(), expectedLog) {
		t.Errorf("expected log to contain %q, got %q", expectedLog, logOutput.String())
	}
}

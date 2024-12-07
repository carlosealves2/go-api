package goapi

import (
	"context"
	"net/http"
	"testing"
)

func TestParamsFromContext(t *testing.T) {
	t.Run("Context contains parameters", func(t *testing.T) {
		params := map[string]string{"key1": "value1", "key2": "value2"}
		ctx := context.WithValue(context.Background(), paramsKey, params)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req = req.WithContext(ctx)

		result := ParamsFromContext(req)
		if result == nil || len(result) != 2 || result["key1"] != "value1" || result["key2"] != "value2" {
			t.Errorf("expected %v, got %v", params, result)
		}
	})

	t.Run("Context does not contain parameters", func(t *testing.T) {
		ctx := context.Background()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req = req.WithContext(ctx)

		result := ParamsFromContext(req)
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})

	// Test case: Context contains a different type
	t.Run("Context contains a different type", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), paramsKey, "not a map")
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req = req.WithContext(ctx)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected a panic when casting to map[string]string")
			}
		}()
		_ = ParamsFromContext(req)
	})
}

package goapi

import "net/http"

func ParamsFromContext(r *http.Request) map[string]string {
	if p := r.Context().Value(paramsKey); p != nil {
		return p.(map[string]string)
	}
	return nil
}

package goapi

import "net/http"

// ParamsFromContext retrieves the route parameters from the HTTP request's context.
// It returns a map of string to string, where the keys are the parameter names and the values are the corresponding parameter values.
// If no route parameters are found in the request's context, it returns nil.
//
// The function uses the provided HTTP request and checks if the request's context contains a value associated with the paramsKey.
// If a value is found, it is assumed to be a map of string to string representing the route parameters.
// The function then returns this map. If no value is found, it returns nil.
func ParamsFromContext(r *http.Request) map[string]string {
	if p := r.Context().Value(paramsKey); p != nil {
		return p.(map[string]string)
	}
	return nil
}

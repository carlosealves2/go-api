package goapi

import (
	"regexp"
	"strings"
)

type route struct {
	method     string
	pattern    *regexp.Regexp
	paramNames []string
	handler    HandlerFunc
}

// parsePattern takes a URL pattern string and returns a compiled regular expression and a slice of parameter names.
// The regular expression matches the URL pattern and extracts the parameter values.
//
// The URL pattern can contain parameter placeholders in the form of ":paramName".
// For example, "/users/:id" will match "/users/123" and extract the "id" parameter value as "123".
//
// The function splits the URL pattern into parts, identifies parameter placeholders, and constructs a regular expression
// that matches the URL pattern. It also extracts the parameter names from the pattern.
//
// The returned regular expression can be used to match incoming URLs and extract parameter values.
// The returned slice of parameter names contains the names of the parameters in the URL pattern.
//
// Example:
//
//	pattern := "/users/:id"
//	regex, paramNames := parsePattern(pattern)
//	fmt.Println(regex.String()) // Output: ^/users/([^/]+)$
//	fmt.Println(paramNames)     // Output: [id]
func parsePattern(pattern string) (*regexp.Regexp, []string) {
	parts := strings.Split(pattern, "/")
	paramNames := make([]string, 0)
	regexParts := make([]string, 0)

	for _, part := range parts {
		if part == "" {
			continue
		}

		if strings.HasPrefix(part, ":") {
			paramName := part[1:]
			paramNames = append(paramNames, paramName)
			regexParts = append(regexParts, "([^/]+)")
		} else {
			regexParts = append(regexParts, part)
		}
	}
	regexPattern := "^/" + strings.Join(regexParts, "/") + "$"
	return regexp.MustCompile(regexPattern), paramNames
}

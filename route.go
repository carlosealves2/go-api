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

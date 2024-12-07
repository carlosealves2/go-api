package goapi

import "testing"

func TestParsePattern(t *testing.T) {
	tests := []struct {
		pattern        string
		expectedReg    string
		expectedParams []string
	}{
		{
			pattern:        "/users/:id",
			expectedReg:    "^/users/([^/]+)$",
			expectedParams: []string{"id"},
		},
		{
			pattern:        "/products/:category/:id",
			expectedReg:    "^/products/([^/]+)/([^/]+)$",
			expectedParams: []string{"category", "id"},
		},
		{
			pattern:        "/static/files",
			expectedReg:    "^/static/files$",
			expectedParams: []string{},
		},
		{
			pattern:        "/:entity/:action/:id",
			expectedReg:    "^/([^/]+)/([^/]+)/([^/]+)$",
			expectedParams: []string{"entity", "action", "id"},
		},
		{
			pattern:        "/",
			expectedReg:    "^/$",
			expectedParams: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.pattern, func(t *testing.T) {
			reg, params := parsePattern(test.pattern)

			if reg.String() != test.expectedReg {
				t.Errorf("expected regex %q, got %q", test.expectedReg, reg.String())
			}

			if len(params) != len(test.expectedParams) {
				t.Errorf("expected %d params, got %d", len(test.expectedParams), len(params))
			}

			for i, param := range params {
				if param != test.expectedParams[i] {
					t.Errorf("expected param %q, got %q", test.expectedParams[i], param)
				}
			}
		})
	}
}

func TestRegexMatching(t *testing.T) {
	pattern := "/users/:id"
	reg, _ := parsePattern(pattern)

	testURL := "/users/123"
	matches := reg.MatchString(testURL)
	if !matches {
		t.Errorf("expected URL %q to match regex %q, but it did not", testURL, reg.String())
	}

	nonMatchingURL := "/users"
	if reg.MatchString(nonMatchingURL) {
		t.Errorf("expected URL %q not to match regex %q, but it did", nonMatchingURL, reg.String())
	}
}

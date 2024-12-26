package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "normalize path",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "cut no-path",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev",
		},
		{
			name:     "long path",
			inputURL: "https://blog.boot.dev/this/is/a/long/path/",
			expected: "blog.boot.dev/this/is/a/long/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

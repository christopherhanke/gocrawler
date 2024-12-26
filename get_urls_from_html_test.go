package main

import "testing"

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no links in HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<h1>No links here!</h1>
				<p>Just some text...</p>
			</body>
		</html>
		`,
			expected: []string{},
		},
		{
			name:     "malformed HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<<>>>> THIS IS NOT HTML
				<a href="https://blog.boot.dev/path">Valid Link</a>
				<a href="/%invalid%path">Invalid Path</a>
			`,
			expected: []string{"https://blog.boot.dev/path"},
		},
		{
			name:     "empty href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="">Empty link</a>
				<a href="    ">Space only link</a>
			</body>
		</html>
		`,
			expected: []string{}, // expecting empty slice since no valid links
		},
		{
			name:     "mixed URL types",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="">Empty link</a>
				<a href="    ">Space only link</a>
				<a href="/">Root path only</a>
				<a href="/valid/path">Valid relative path</a>
				<a href="https://blog.boot.dev">Base URL only</a>
				<a href="https://blog.boot.dev/valid/path">Valid absolute path</a>
			</body>
		</html>
		`,
			expected: []string{
				"https://blog.boot.dev/valid/path",
				"https://blog.boot.dev/valid/path",
			},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if len(actual) != len(tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URLs: %v, actual URLs: %v", i, tc.name, len(tc.expected), len(actual))
				return
			}
			for j := range actual {
				if actual[j] != tc.expected[j] {
					t.Errorf("Test %v - %s FAIL: expected: %s, actual: %s", i, tc.name, tc.expected[j], actual[j])
				}
			}
		})
	}
}

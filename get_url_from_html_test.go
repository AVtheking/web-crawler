package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetUrlFromHtml(t *testing.T) {
	tests := []struct {
		name          string
		inputUrl      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:      "GetUrlFromHtml",
			inputUrl:  "https://blog.boot.dev/path",
			inputBody: "<a href='https://blog.boot.dev/path/'>Link</a>",
			expected:  []string{"https://blog.boot.dev/path/"},
		},
		{
			name:     "absolute and relative URLs",
			inputUrl: "https://blog.boot.dev",
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
			name:      "no links found",
			inputUrl:  "https://blog.boot.dev",
			inputBody: "<html><body><h1>Hello, World!</h1></body></html>",
			expected:  nil,
		},
		{
			name:     "invalid url",
			inputUrl: "https://example.com",
			inputBody: `
		<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
		`,
			expected: nil,
		}, {
			name:     "invalid base url",
			inputUrl: `:\\invalidBaseURL`,
			inputBody: `
<html>
	<body>
		<a href="/path">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected:      nil,
			errorContains: "Error parsing base url:",
		},
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualUrls, err := GetUrlsFromHtml(test.inputBody, test.inputUrl)
			if err != nil && !strings.Contains(err.Error(), test.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, test.name, err)
				return
			} else if err != nil && test.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, test.name, err)
				return
			} else if err == nil && test.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%s', but got none", i, test.name, test.errorContains)
				return
			}
			if !reflect.DeepEqual(actualUrls, test.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected %v, but got %v", i, test.name, test.expected, actualUrls)
			}
		})
	}
}

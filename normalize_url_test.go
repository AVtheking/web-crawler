package main

import "testing"

func TestNormalizeUrl(t *testing.T) {
	tests := []struct {
		name        string
		inputUrl    string
		expectedUrl string
	}{

		{
			name:        "NormalizeUrl",
			inputUrl:    "https://blog.boot.dev/path",
			expectedUrl: "blog.boot.dev/path",
		},
		{
			name:        "NormalizeUrl with trailing slash",
			inputUrl:    "https://blog.boot.dev/path/",
			expectedUrl: "blog.boot.dev/path",
		},
		{
			name:        "NormalizeUrl with http",
			inputUrl:    "http://blog.boot.dev/path/",
			expectedUrl: "blog.boot.dev/path",
		},
		{
			name:        "url with capital letters",
			inputUrl:    "https://BLOG.BOOT.DEV/path/",
			expectedUrl: "blog.boot.dev/path",
		},
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualUrl, err := NormalizeUrl(test.inputUrl)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, test.name, err)
				return
			}
			if actualUrl != test.expectedUrl {
				t.Errorf("Test %v - '%s' FAIL: expected %s, but got %s", i, test.name, test.expectedUrl, actualUrl)
			}
		})
	}

}

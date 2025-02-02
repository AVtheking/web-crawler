package main

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeUrl(urlString string) (string, error) {
	//take the url and normalize it
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return "", fmt.Errorf("failed to parse url: %v", err)
	}

	fullPath := parsedUrl.Host + parsedUrl.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")
	return fullPath, nil
}

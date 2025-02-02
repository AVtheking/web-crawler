package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawUrl string) (string, error) {
	//get the html of the page
	response, err := http.Get(rawUrl)
	if err != nil {
		return "", fmt.Errorf("error getting html: %w", err)
	}
	//check if the response is ok
	if response.StatusCode >= 400 {
		return "", fmt.Errorf("error getting html: %s", response.Status)
	}
	//check if the content type is html
	contentType := response.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("error getting content type: Content-Type header is missing")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading html: %w", err)
	}
	return string(body), nil
}

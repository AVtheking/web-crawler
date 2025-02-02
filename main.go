package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	//get the args from the command line
	args := os.Args
	argsWithoutProgramName := args[1:]
	if len(argsWithoutProgramName) < 1 {
		fmt.Println("no website provided ")
		os.Exit(1)
	}
	if len(argsWithoutProgramName) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseUrl := argsWithoutProgramName[0]
	fmt.Println("starting crawl of ", baseUrl)
	html, err := getHTML(baseUrl)
	if err != nil {
		fmt.Println("error getting html: ", err)
		os.Exit(1)
	}
	fmt.Println(html)

}

func getHTML(rawUrl string) (string, error) {
	response, err := http.Get(rawUrl)
	if err != nil {
		return "", fmt.Errorf("error getting html: %w", err)
	}
	if response.StatusCode >= 400 {
		return "", fmt.Errorf("error getting html: %s", response.Status)
	}
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

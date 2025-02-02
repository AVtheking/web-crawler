package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
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

	pages := make(map[string]int)
	pages = crawlPage(baseUrl, baseUrl, pages)
	for url, count := range pages {
		fmt.Printf("URL: %s, Count: %d\n", url, count)
	}

}

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

// function to recursively crawl the website
func crawlPage(rawBaseURL string, rawCurrentURL string, pages map[string]int) map[string]int {
	// if current url is on different domain then return pages
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("error parsing current url: ", err)
		return pages
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("error parsing base url: ", err)
		return pages
	}

	if currentURL.Hostname() != baseURL.Hostname() {
		return pages
	}
	//normalize the url
	normalizedURL, err := NormalizeUrl(rawCurrentURL)
	if err != nil {
		fmt.Println("error normalizing url: ", err)
		return pages
	}

	//if the page is already visisted then  increase the count
	if _, ok := pages[normalizedURL]; ok {
		pages[normalizedURL]++
		return pages
	} else {
		pages[normalizedURL] = 1
	}

	//get the html of the page
	fmt.Println("crawling: ", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("error getting html: ", err)
		return pages
	}
	// fmt.Println("html of the page %s: ", rawCurrentURL, html)
	//get the urls from the html
	urls, err := GetUrlsFromHtml(html, rawBaseURL)
	if err != nil {
		fmt.Println("error getting urls from html: ", err)
		return pages
	}
	// fmt.Println("URLS found on the page %s: ", rawCurrentURL, urls)
	//for each url found on the page, crawl the page
	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
	return pages

}

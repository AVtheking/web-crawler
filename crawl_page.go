package main

import (
	"fmt"
	"net/url"
)

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

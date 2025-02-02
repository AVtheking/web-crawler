package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) map[string]int {
	// if current url is on different domain then return pages
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("error parsing current url: ", err)
		return cfg.pages
	}
	baseURL, err := url.Parse(cfg.baseURL.String())
	if err != nil {
		fmt.Println("error parsing base url: ", err)
		return cfg.pages
	}

	if currentURL.Hostname() != baseURL.Hostname() {
		return cfg.pages
	}
	//normalize the url
	normalizedURL, err := NormalizeUrl(rawCurrentURL)
	if err != nil {
		fmt.Println("error normalizing url: ", err)
		return cfg.pages
	}

	//if the page is already visisted then  increase the count
	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.pages[normalizedURL]++
		return cfg.pages
	} else {
		cfg.pages[normalizedURL] = 1
	}

	//get the html of the page
	fmt.Println("crawling: ", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("error getting html: ", err)
		return cfg.pages
	}
	// fmt.Println("html of the page %s: ", rawCurrentURL, html)
	//get the urls from the html
	urls, err := GetUrlsFromHtml(html, cfg.baseURL.String())
	if err != nil {
		fmt.Println("error getting urls from html: ", err)
		return cfg.pages
	}
	// fmt.Println("URLS found on the page %s: ", rawCurrentURL, urls)
	//for each url found on the page, crawl the page
	for _, url := range urls {
		cfg.crawlPage(url)
	}
	return cfg.pages

}

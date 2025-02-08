package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	//acting as a buffered channel to control the concurrency
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		//releasing the buffer
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	//check if the max pages is reached
	if cfg.pagesLength() >= cfg.maxPages {
		return
	}

	// if current url is on different domain then return pages
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("error parsing current url: ", err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}
	//normalize the url
	normalizedURL, err := NormalizeUrl(rawCurrentURL)
	if err != nil {
		fmt.Println("error normalizing url: ", err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	//get the html of the page
	fmt.Println("crawling: ", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("error getting html: ", err)
		return
	}
	// fmt.Println("html of the page %s: ", rawCurrentURL, html)
	//get the urls from the html
	urls, err := GetUrlsFromHtml(html, cfg.baseURL.String())
	if err != nil {
		fmt.Println("error getting urls from html: ", err)
		return
	}
	// fmt.Println("URLS found on the page %s: ", rawCurrentURL, urls)
	//for each url found on the page, crawl the page
	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetUrlFromHtml(htmlBody, rawBaseUrl string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("Error parsing html: %w", err)

	}

	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {

		return nil, fmt.Errorf("Error parsing base url: %w", err)
	}

	var urls []string
	var f func(*html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					parsedUrl, err := url.Parse(attr.Val)
					if err != nil {
						continue
					}
					absoluteUrl := baseUrl.ResolveReference(parsedUrl)
					urls = append(urls, absoluteUrl.String())
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return urls, nil
}

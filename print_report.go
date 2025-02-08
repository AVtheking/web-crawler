package main

import (
	"fmt"
	"sort"
)

type Page struct {
	URL   string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("========================================================")
	fmt.Println("REPORT for", baseURL)
	fmt.Println("========================================================")
	for _, page := range sortPages(pages) {
		fmt.Printf("Found %d internal links to %s\n", page.count, page.URL)
	}
}

// sort pages by count into a slice of structs
func sortPages(pages map[string]int) []Page {
	pageList := make([]Page, 0, len(pages))
	for url, count := range pages {

		pageList = append(pageList, Page{URL: url, count: count})
	}
	sort.Slice(pageList, func(i, j int) bool {
		return pageList[i].count > pageList[j].count
	})
	return pageList
}

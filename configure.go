package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

// addPageVisit adds a page visit to the config with a mutex lock
func (cfg *config) addPageVisit(url string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, ok := cfg.pages[url]; ok {
		cfg.pages[url]++
		return false
	}
	cfg.pages[url] = 1
	return true
}

func configure(rawBaseURL string, maxConcurrency int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base url: %w", err)
	}
	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}

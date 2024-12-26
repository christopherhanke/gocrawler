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

func configure(rawBaseURL string, maxConcurrency int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("could not parse URL: %s, %v", rawBaseURL, err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}

func (cfg *config) addPageVisit(normilzedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[normilzedURL]; ok {
		cfg.pages[normilzedURL]++
		return false
	}

	cfg.pages[normilzedURL] = 1
	return true
}

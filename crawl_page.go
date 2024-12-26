package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	// setup concurrent crawling
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	// parse rawURL to URL
	currURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("error parsing current url. state: %s, %v", rawCurrentURL, len(cfg.pages))
		return
	}

	// chekc if currentURL is on baseURL host
	if cfg.baseURL.Hostname() != currURL.Hostname() {
		return
	}

	// normalize current URL and count visits in pages
	normalCurr, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("error normalizing: %v", err)
		return
	}

	// increment if visited
	firstTime := cfg.addPageVisit(normalCurr)
	if !firstTime {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	// get HTML body and get the URL from current site
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("error getting HTML from %s: %v", rawCurrentURL, err)
		return
	}
	links, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		log.Printf("error getting URLs from HTML, current: %s error: %v", rawCurrentURL, err)
		return
	}

	// recursively call crawlPage with new links
	for _, link := range links {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}

}

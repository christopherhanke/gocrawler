package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// parse rawURL to URL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Printf("error parsing base url. state: %s, %s, %v", rawBaseURL, rawCurrentURL, len(pages))
		return
	}
	currURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("error parsing current url. state: %s, %s, %v", rawBaseURL, rawCurrentURL, len(pages))
		return
	}

	// chekc if currentURL is on baseURL host
	if baseURL.Hostname() != currURL.Hostname() {
		return
	}

	// normalize current URL and count visits in pages
	normalCurr, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("error normalizing: %v", err)
		return
	}

	// increment if visited
	if _, ok := pages[normalCurr]; ok {
		pages[normalCurr]++
		return
	}

	// mark as visited
	pages[normalCurr] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

	// get HTML body and get the URL from current site
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("error getting HTML from %s: %v", rawCurrentURL, err)
		return
	}
	links, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		log.Printf("error getting URLs from HTML, current: %s error: %v", rawCurrentURL, err)
		return
	}

	// recursively call crawlPage with new links
	for _, link := range links {
		crawlPage(rawBaseURL, link, pages)
	}

}

package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	// parse baseURL to handle only host
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	// read the HTML body and parse it to html nodes
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	// visit nodes and write links in result
	var result []string
	visitNode(doc, &result)

	// read results and handle relative urls
	validURLs := make([]string, 0, len(result))
	for _, urlString := range result {
		urlString = strings.TrimSpace(urlString)
		if urlString == "" {
			continue
		}
		parsedURL, err := url.Parse(urlString)
		if err != nil {
			continue
		}

		if parsedURL.Hostname() == "" {
			parsedURL = baseURL.ResolveReference(parsedURL)
		}
		if parsedURL.Path == "" || parsedURL.Path == "/" {
			continue
		}
		validURLs = append(validURLs, parsedURL.String())
	}
	return validURLs, nil
}

// helper function to traverse html node tree and write all links from <a>-Nodes in the urls []string.
func visitNode(n *html.Node, urls *[]string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				*urls = append(*urls, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visitNode(c, urls)
	}
}

package main

import (
	"net/url"
	"strings"
)

func normalizeURL(s string) (string, error) {
	parsedURL, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	return parsedURL.Hostname() + strings.TrimRight(parsedURL.Path, "/"), nil
}

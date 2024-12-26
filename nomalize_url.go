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
	fullPath := parsedURL.Host + parsedURL.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}

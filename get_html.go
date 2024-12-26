package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 500 && resp.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("response error: %v", resp.StatusCode)
	}

	if !strings.Contains(resp.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("response content type mismatched: %v", resp.Header.Get("content-type"))
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

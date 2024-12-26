package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("not enough arguments")
		fmt.Println("usage: go run . 'urlWebsite' 'maxConcurrency' 'maxPages'")
		return
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		fmt.Println("usage: go run . 'urlWebsite' 'maxConcurrency' 'maxPages'")
		return
	}

	rawBaseURL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("maxConcurrency not valid: %v\n", err)
		return
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("maxPages not valid: %v\n", err)
		return
	}

	// setup configure
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("error creating configure: %v", err)
		return
	}
	// starting crawl
	fmt.Printf("starting crawl of: %s\n", args[0])
	cfg.wg.Add(1)
	go cfg.crawlPage(args[0])
	cfg.wg.Wait()

	// printing report
	for k := range cfg.pages {
		fmt.Printf("Page: %s, visited: %d\n", k, cfg.pages[k])
	}

}

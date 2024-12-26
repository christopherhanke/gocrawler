package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		return
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		return
	}

	// setup configure
	cfg, err := configure(args[0], 3)
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

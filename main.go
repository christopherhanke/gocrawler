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
	fmt.Printf("starting crawl of: %s\n", args[0])
	pages := make(map[string]int)
	crawlPage(args[0], args[0], pages)

	for k := range pages {
		fmt.Printf("Page: %s, visited: %d\n", k, pages[k])
	}

}

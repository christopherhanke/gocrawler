package main

import (
	"fmt"
	"sort"
)

type PageCount struct {
	URL   string
	Count int
}

func printReport(pages map[string]int, baseURL string) {
	//fmt.Printf("Len: %d", len(pages))
	sortedPages := sortReport(pages)

	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=============================")
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}

}

func sortReport(pages map[string]int) []PageCount {
	//fmt.Printf("sortReport, Len: %d", len(pages))
	var pagesCount []PageCount
	for url, count := range pages {
		pagesCount = append(pagesCount, PageCount{
			URL:   url,
			Count: count,
		})
	}
	//fmt.Printf("sortReport, len new: %d", len(pagesCount))
	sort.Slice(pagesCount, func(i, j int) bool {
		if pagesCount[i].Count > pagesCount[j].Count {
			return true
		}
		if pagesCount[i].Count == pagesCount[j].Count {
			return pagesCount[i].URL < pagesCount[j].URL
		}
		return false
	})
	return pagesCount
}

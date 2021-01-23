package main

import (
	"fmt"
	"log"
	"os"

	links "golang-study/book/ch5_functions/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)

	go func() {
		workList <- os.Args[1:]
	}()

	// crawl the web concurrently
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(l string) { // too parallel, socket error: too many open files
					workList <- crawl(l)
				}(link)
			}
		}
	}
}

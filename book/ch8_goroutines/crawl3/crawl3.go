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

// 未实现退出逻辑
func main() {
	workList := make(chan []string)
	unseenLinks := make(chan string)

	go func() { workList <- os.Args[1:] }()

	// create 20 crawler goroutines to fetch each unseen link
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { workList <- foundLinks }()
			}
		}()
	}

	// the main goroutines send unseen links to the crawlers
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	links "golang-study/book/ch5_functions/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{} //acquire a token to do job
	list, err := links.Extract(url)
	<-tokens //release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)
	var n int //number of pending sends to workList

	n++
	go func() { workList <- os.Args[1:] }()

	// crawl the web concurrently
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++ // add a new send
				go func(l string) {
					workList <- crawl(l)
				}(link)
			}
		}
	}
}

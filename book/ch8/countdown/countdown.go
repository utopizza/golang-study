package main

import (
	"fmt"
	"time"
)

func main() {
	abort := make(chan struct{})
	ticker := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker:
			// do nothing
		case <-abort:
			fmt.Println("launch abort!")
			return
		}
	}
	//launch()
}

package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out) // close channel to notiy receivers
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out) // close channel to notiy receivers
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

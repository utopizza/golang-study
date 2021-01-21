package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals) // close channel to notify receivers
	}()

	//Squarer
	go func() {
		// range loop to receive from channel, exit while channel is closed
		for x := range naturals {
			squares <- x * x
		}
		close(squares) // close channel to notify receivers
	}()

	//Printer
	for x := range squares {
		fmt.Println(x)
	}
}

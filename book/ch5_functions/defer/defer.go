package ch5

import "fmt"

func double(x int) int {
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }() // executed after return
	return double(x)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panic if x==0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	fmt.Println(triple(4)) //12

	// when a panic occurs,
	// all deferred functions are run
	// in reverse order
	f(3)
}

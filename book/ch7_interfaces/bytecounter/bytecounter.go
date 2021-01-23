package ch7

import "fmt"

// ByteCounter 字节计数器
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) //5

	c = 0
	fmt.Fprintf(&c, "hello, %s", "Dolly")
	fmt.Println(c) //12
}

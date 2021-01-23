package ch7

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// containsAll reports whether x contains the elements of y, in order
func containsAll(x, y []string) bool {
	for len(y) < len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xml select:%v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) //push
		case xml.EndElement:
			stack = stack[:len(stack)-1] //pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

package ch8

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show progress")

// control the count of goroutines
var tokens = make(chan struct{}, 20)

func dirEntries(dir string) []os.FileInfo {
	tokens <- struct{}{}        //acquire token
	defer func() { <-tokens }() //release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	return entries
}

func walkDir(dir string, fileSize chan<- int64, n *sync.WaitGroup) {
	defer n.Done()
	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, fileSize, n) //concurrently
		} else {
			fileSize <- entry.Size()
		}
	}
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// traverse the file tree
	fileSize := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, fileSize, &n) //concurrently
	}

	// closer
	go func() {
		n.Wait()
		close(fileSize)
	}()

	// print the sum periodically
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nFiles, nBytes int64
loop:
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				break loop
			}
			nFiles++
			nBytes += size
		case <-tick:
			fmt.Printf("%d files, %.1f GB\n", nFiles, float64(nBytes)/1e9)
		}
	}
	fmt.Printf("%d files, %.1f GB\n", nFiles, float64(nBytes)/1e9)
}

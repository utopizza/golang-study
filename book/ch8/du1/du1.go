package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 读出目录下的列表
func dirEntries(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	return entries
}

// 深度优先遍历目录，若遇到文件则将文件大小发送到通道
func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileSize)
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
	go func() {
		for _, root := range roots {
			walkDir(root, fileSize)
		}
		close(fileSize)
	}()

	// sum up the filesize from channel
	var nFiles, nBytes int64
	for size := range fileSize {
		nFiles++
		nBytes += size
	}

	fmt.Printf("%d files, %.1f GB\n", nFiles, float64(nBytes)/1e9)
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func TestMemo_Get_Concurrently(t *testing.T) {
	cache := NewMemo(httpGetBody)
	var wg sync.WaitGroup
	urls := []string{
		"http://www.baidu.com", "http://www.baidu.com",
		"http://www.github.com", "http://www.github.com",
		"http://www.qq.com", "http://www.qq.com",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := cache.Get(url)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}

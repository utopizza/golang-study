package ch5

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// WaitForServer 失败重试到成功，或者到达ddl为止
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil //success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) //exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

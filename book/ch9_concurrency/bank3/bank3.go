package main

import "sync"

var (
	mu      sync.Mutex // 用mutex实现互斥访问
	balance int
)

func deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func getBalance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

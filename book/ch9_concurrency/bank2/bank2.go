package main

var (
	sema    = make(chan struct{}, 1) // 用channel实现互斥访问
	balance int
)

func deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

func getBalance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

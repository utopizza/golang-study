package main

var deposits = make(chan int)
var balances = make(chan int)

func deposit(amount int) {
	deposits <- amount
}

func balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func main() {
	go teller()
}

package main

type cake struct {
	state string
}

func baker(cooked chan<- *cake) {
	for {
		cake := new(cake)
		cake.state = "cooked"
		cooked <- cake
	}
}

func icer(iced chan<- *cake, cooked <-chan *cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake
	}
}

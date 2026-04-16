package main

func main() {
	aapl := NewStock("AAPL", 150)
	tsla := NewStock("TSLA", 200)

	alice := NewInvestor("Alice", &AggresiveStrategy{})
	bob := NewInvestor("Bob", &ConservativeStrategy{})

	aapl.Register(alice)
	aapl.Register(bob)

	aapl.SetPrice(180)
	aapl.SetPrice(90)

	alice.SetStrategy(&ConservativeStrategy{})
	aapl.SetPrice(85)

	tsla.Register(alice)
	tsla.SetPrice(300)
}

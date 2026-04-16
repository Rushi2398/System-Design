package main

import "fmt"

type Investor struct {
	name     string
	strategy TradingStrategy
}

func NewInvestor(name string, strategy TradingStrategy) *Investor {
	return &Investor{name: name, strategy: strategy}
}

func (i *Investor) Update(stock string, price float64) {
	fmt.Printf("%s received update for %s → ", i.name, stock)
	i.strategy.DecideAction(stock, price)
}

func (i *Investor) SetStrategy(strategy TradingStrategy) {
	fmt.Printf("\n%s switched strategy.\n", i.name)
	i.strategy = strategy
}

package main

import "fmt"

type AggresiveStrategy struct{}
type ConservativeStrategy struct{}

func (a *AggresiveStrategy) DecideAction(stock string, price float64) {
	fmt.Printf("[Aggressive] %s price: %.2f → Action: BUY immediately!\n", stock, price)
}

func (c *ConservativeStrategy) DecideAction(stock string, price float64) {
	if price < 100 {
		fmt.Printf("[Conservative] %s price: %.2f → Action: BUY (undervalued)\n", stock, price)
	} else {
		fmt.Printf("[Conservative] %s price: %.2f → Action: HOLD\n", stock, price)
	}
}

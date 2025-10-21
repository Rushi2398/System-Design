package main

type TradingStrategy interface {
	DecideAction(stock string, price float64)
}

package main

type Observer interface {
	Update(stock string, price float64)
}

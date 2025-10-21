package main

import "fmt"

type Stock struct {
	stock     string
	price     float64
	observers []Observer
}

func NewStock(name string, price float64) *Stock {
	return &Stock{stock: name, price: price}
}

func (s *Stock) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Stock) Unregister(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Stock) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.stock, s.price)
	}
}

func (s *Stock) SetPrice(price float64) {
	fmt.Printf("\n[Stock] %s new price: %.2f\n", s.stock, price)
	s.price = price
	s.Notify()
}

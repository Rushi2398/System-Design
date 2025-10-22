package main

// Base Concrete Beverage
type Espresso struct{}

func (Espresso) Cost() float64 {
	return 1.99
}

func (Espresso) Description() string {
	return "Espresso"
}

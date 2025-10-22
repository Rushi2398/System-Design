package main

type Milk struct {
	BeverageDecorator
}

func NewMilk(b Beverage) Beverage {
	return Milk{BeverageDecorator{Beverage: b}}
}

func (m Milk) Cost() float64 {
	return m.Beverage.Cost() + 0.20
}

func (m Milk) Description() string {
	return m.Beverage.Description() + ", Milk"
}

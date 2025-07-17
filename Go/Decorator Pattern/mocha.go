package main

type Mocha struct {
	BeverageDecorator
}

func NewMocha(b Beverage) Beverage {
	return Mocha{BeverageDecorator{Beverage: b}}
}

func (m Mocha) Cost() float64 {
	return m.Beverage.Cost() + 0.30
}

func (m Mocha) Description() string {
	return m.Beverage.Description() + ", Mocha"
}

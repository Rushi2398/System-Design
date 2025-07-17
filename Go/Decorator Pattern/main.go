package main

import "fmt"

func main() {
	var drink Beverage = Espresso{}
	fmt.Printf("%s: $%.2f\n", drink.Description(), drink.Cost())

	drink = NewMilk(drink)
	fmt.Printf("%s: $%.2f\n", drink.Description(), drink.Cost())

	drink = NewMocha(drink)
	fmt.Printf("%s: $%.2f\n", drink.Description(), drink.Cost())
}

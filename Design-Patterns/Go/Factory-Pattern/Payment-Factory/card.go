package main

import "fmt"

type Card struct {
	Number string
}

func (c Card) Pay(amount int) string {
	return fmt.Sprintf("Card(%s) paid %d", c.Number, amount)
}

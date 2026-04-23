package main

import "fmt"

type UPI struct {
	ID string
}

func (u UPI) Pay(amount int) string {
	return fmt.Sprintf("UPI(%s) paid %d", u.ID, amount)
}

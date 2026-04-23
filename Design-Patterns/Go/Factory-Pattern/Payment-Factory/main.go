package main

import "fmt"

func init() {
	Register("upi", func(args ...any) Payment {
		return UPI{ID: args[0].(string)}
	})

	Register("card", func(args ...any) Payment {
		return Card{Number: args[0].(string)}
	})
}

// 7. Main

func main() {
	p1 := Create("upi", "user@upi")
	p2 := Create("card", "1234-xxxx")

	fmt.Println(p1.Pay(1000))
	fmt.Println(p2.Pay(2000))
}

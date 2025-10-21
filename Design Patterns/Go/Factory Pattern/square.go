package main

import "fmt"

type Square struct{}

func (s Square) Draw() {
	fmt.Println("Drawing a Square")
}

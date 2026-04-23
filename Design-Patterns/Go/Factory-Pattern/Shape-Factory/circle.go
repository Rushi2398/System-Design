package main

import "fmt"

type Circle struct{}

func (c Circle) Draw() {
	fmt.Println("Drawing a Circle")
}

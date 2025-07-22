package main

import "fmt"

type Rectangle struct{}

func (r Rectangle) Draw() {
	fmt.Println("Drawing a Rectangle")
}

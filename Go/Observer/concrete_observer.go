package main

import "fmt"

type ConcreteObserver struct {
	Name string
}

func (co *ConcreteObserver) Update(data string) {
	fmt.Printf("%s received: %s\n", co.Name, data)
}

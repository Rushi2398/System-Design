package main

import "fmt"

type Navigator struct {
	navigation NavigationStrategy
}

func (n *Navigator) SetStrategy(navigation NavigationStrategy) {
	n.navigation = navigation
}

func (n *Navigator) Navigate(source, destination string) {
	if n.navigation == nil {
		fmt.Println("No Navigation Strategy is set!")
		return
	}
	n.navigation.Navigate(source, destination)
}

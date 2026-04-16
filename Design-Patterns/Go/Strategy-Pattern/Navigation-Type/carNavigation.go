package main

import "fmt"

type RoadNavigation struct{}

func (r *RoadNavigation) Navigate(source, destination string) {
	fmt.Printf("Navigating from %s to %s via Car.\n", source, destination)
}

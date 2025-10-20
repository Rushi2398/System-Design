package main

import "fmt"

type BusNavigation struct{}

func (r *BusNavigation) Navigate(source, destination string) {
	fmt.Printf("Navigating from %s to %s via Bus.\n", source, destination)
}

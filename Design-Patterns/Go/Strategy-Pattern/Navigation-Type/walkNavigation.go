package main

import "fmt"

type WalkNavigation struct{}

func (r *WalkNavigation) Navigate(source, destination string) {
	fmt.Printf("Navigating from %s to %s via Road.\n", source, destination)
}

package main

import "fmt"

type SpecialDrive struct{}

func (sd SpecialDrive) Drive() {
	fmt.Println("Special Drive Capability")
}

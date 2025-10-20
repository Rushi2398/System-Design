package main

import "fmt"

type NormalDrive struct{}

func (nd NormalDrive) Drive() {
	fmt.Println("Normal Drive Capability")
}

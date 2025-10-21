package main

import "fmt"

func main() {
	fmt.Println("== Windows UI ==")
	RenderUI(&WindowsFactory{})

	fmt.Println("\n== MacOS UI ==")
	RenderUI(&MacFactory{})
}

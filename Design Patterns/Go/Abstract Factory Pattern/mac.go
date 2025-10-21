package main

import "fmt"

type MacButton struct{}

func (b *MacButton) Click() {
	fmt.Println("Mac Button Clicked!")
}

type MacWindow struct{}

func (w *MacWindow) Open() {
	fmt.Println("Mac Window Opened!")
}

type MacFactory struct{}

func (f *MacFactory) CreateButton() Button { return &MacButton{} }
func (f *MacFactory) CreateWindow() Window { return &MacWindow{} }

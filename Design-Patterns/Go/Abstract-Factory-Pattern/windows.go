package main

import "fmt"

type WindowsButton struct{}

func (b *WindowsButton) Click() {
	fmt.Println("Windows Button Clicked!")
}

type WindowsWindow struct{}

func (w *WindowsWindow) Open() {
	fmt.Println("Windows Window Opened!")
}

type WindowsFactory struct{}

func (f *WindowsFactory) CreateButton() Button { return &WindowsButton{} }
func (f *WindowsFactory) CreateWindow() Window { return &WindowsWindow{} }

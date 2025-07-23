package main

// Product Interface
type Button interface {
	Click()
}

type Window interface {
	Open()
}

// Abstract Factory Interface
type GUIFactory interface {
	CreateButton() Button
	CreateWindow() Window
}

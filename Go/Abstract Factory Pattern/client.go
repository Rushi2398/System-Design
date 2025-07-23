package main

func RenderUI(factory GUIFactory) {
	btn := factory.CreateButton()
	window := factory.CreateWindow()
	btn.Click()
	window.Open()
}

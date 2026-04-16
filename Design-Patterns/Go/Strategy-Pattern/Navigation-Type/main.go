package main

func main() {
	navigator := &Navigator{}

	navigator.SetStrategy(&RoadNavigation{})
	navigator.Navigate("Home", "Office")

	navigator.SetStrategy(&WalkNavigation{})
	navigator.Navigate("Park", "Cafe")

	navigator.SetStrategy(&BusNavigation{})
	navigator.Navigate("Airport", "Hotel")
}

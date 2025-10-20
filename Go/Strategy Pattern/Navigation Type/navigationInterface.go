package main

type NavigationStrategy interface {
	Navigate(source, destination string)
}

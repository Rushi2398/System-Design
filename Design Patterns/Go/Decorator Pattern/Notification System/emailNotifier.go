package main

import "fmt"

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("Sending Email: ", message)
}

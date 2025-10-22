package main

import "fmt"

type SMSNotifier struct {
	*NotifierDecorator
}

func (s *SMSNotifier) Send(message string) {
	s.Notifier.Send(message)
	fmt.Println("Sending SMS: ", message)
}

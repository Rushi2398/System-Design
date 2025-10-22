package main

import "fmt"

type PushNotifier struct {
	*NotifierDecorator
}

func (p *PushNotifier) Send(message string) {
	p.Notifier.Send(message)
	fmt.Println("Sending Push Notification: ", message)
}

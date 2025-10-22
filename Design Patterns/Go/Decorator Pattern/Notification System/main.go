package main

func main() {
	var notifier Notifier = &EmailNotifier{}
	notifier = &PushNotifier{&NotifierDecorator{notifier}}
	notifier = &SMSNotifier{&NotifierDecorator{notifier}}

	notifier.Send("Hello, Rushikesh")
}

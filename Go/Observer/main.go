package main

func main() {
	sub := &ConcreteSubject{}

	obs1 := &ConcreteObserver{Name: "Observer 1"}
	obs2 := &ConcreteObserver{Name: "Observer 2"}

	sub.Register(obs1)
	sub.Register(obs2)

	sub.Notify("First Notification")

	sub.Unregister(obs1)

	sub.Notify("Second Notification")
}

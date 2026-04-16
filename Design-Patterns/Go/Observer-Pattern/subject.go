package main

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(data string)
	GetState(data string)
}

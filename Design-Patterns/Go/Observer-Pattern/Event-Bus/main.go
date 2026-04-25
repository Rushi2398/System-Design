package main

import (
	"fmt"
	"sync"
)

type EventBus struct {
	subscribers map[string][]func(any)
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(any)),
	}
}

func (eb *EventBus) Subscribe(event string, fn func(any)) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.subscribers[event] = append(eb.subscribers[event], fn)
}

func (eb *EventBus) Publish(event string, data any) {
	eb.mu.RLock()
	listeners := eb.subscribers[event]
	eb.mu.RUnlock()

	for _, fn := range listeners {
		fn(data)
	}
}

func main() {
	bus := NewEventBus()

	bus.Subscribe("order_created", func(data any) {
		order := data.(map[string]string)
		fmt.Println("Email sent for ", order["id"])
	})

	bus.Subscribe("order_created", func(data any) {
		order := data.(map[string]string)
		fmt.Println("Inventory updated for", order["id"])
	})

	bus.Publish("order_created", map[string]string{"id": "ORD123"})
}

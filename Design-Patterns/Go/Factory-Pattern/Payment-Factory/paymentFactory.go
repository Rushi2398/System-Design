package main

import "sync"

var (
	registry = make(map[string]func(...any) Payment)
	mu       sync.RWMutex
)

// 3. Register function
func Register(name string, creator func(...any) Payment) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := registry[name]; exists {
		panic("type already registered: " + name)
	}

	registry[name] = creator
}

// 4. Factory
func Create(name string, args ...any) Payment {
	mu.RLock()
	creator, exists := registry[name]
	mu.RUnlock()

	if !exists {
		panic("unknown type: " + name)
	}

	return creator(args...)
}

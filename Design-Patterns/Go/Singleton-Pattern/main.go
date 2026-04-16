package main

import (
	"fmt"
	"sync"
)

type Config struct {
	Port int
}

var (
	instance *Config
	once     sync.Once
)

func GetInstance() *Config {
	once.Do(func() {
		fmt.Println("Creating Singleton Instance...")
		instance = &Config{Port: 8080}
	})
	return instance
}

func main() {
	c1 := GetInstance()
	c2 := GetInstance()

	fmt.Println(c1 == c2)
	fmt.Println(c1.Port)
}

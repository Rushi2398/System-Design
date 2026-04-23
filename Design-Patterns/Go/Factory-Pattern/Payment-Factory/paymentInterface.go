package main

type Payment interface {
	Pay(amt int) string
}

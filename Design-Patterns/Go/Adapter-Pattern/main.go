package main

import (
	"fmt"
	"sync"
)

type Payment interface {
	Pay(amount int) string
}

type Stripe struct{}

func (s Stripe) MakePayment(amount int) string {
	return fmt.Sprintf("Stripe charged %d", amount)
}

type PayPal struct{}

func (p PayPal) SendPayment(amount int) string {
	return fmt.Sprintf("PayPal processed %d", amount)
}

type StripeAdapter struct {
	stripe Stripe
}

func (s StripeAdapter) Pay(amount int) string {
	return s.stripe.MakePayment(amount)
}

type PayPalAdapter struct {
	paypal PayPal
}

func (p PayPalAdapter) Pay(amount int) string {
	return p.paypal.SendPayment(amount)
}

var (
	registry = make(map[string]func() Payment)
	mu       sync.RWMutex
)

func Register(name string, creator func() Payment) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := registry[name]; exists {
		panic("Already registered: " + name)
	}
	registry[name] = creator
}

func Create(name string) Payment {
	mu.RLock()
	creator, exists := registry[name]
	mu.RUnlock()
	if !exists {
		panic("unknown provider: " + name)
	}
	return creator()
}

func init() {
	Register("stripe", func() Payment {
		return StripeAdapter{stripe: Stripe{}}
	})

	Register("paypal", func() Payment {
		return PayPalAdapter{paypal: PayPal{}}
	})
}

func main() {
	provider := "paypal"
	payment := Create(provider)
	fmt.Println(payment.Pay(1000))
}

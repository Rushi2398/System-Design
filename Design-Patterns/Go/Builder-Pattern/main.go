package main

import (
	"fmt"
	"sync"
)

// // 1. External APIs
type Stripe struct{}

func (s Stripe) CreateCharge(amount int) map[string]interface{} {
	return map[string]interface{}{"id": "txn_1", "amount": amount}
}

type PayPal struct{}

func (p PayPal) SendPayment(amount int) map[string]interface{} {
	return map[string]interface{}{"txnId": "txn_2", "total": amount}
}

// // 2. Adapter Interface
type Adapter interface {
	Pay(amount int) map[string]interface{}
}

type StripeAdapter struct {
	stripe Stripe
}

func (s StripeAdapter) Pay(amount int) map[string]interface{} {
	res := s.stripe.CreateCharge(amount)
	return map[string]interface{}{
		"id":     res["id"],
		"amount": res["amount"],
	}
}

type PayPalAdapter struct {
	paypal PayPal
}

func (p PayPalAdapter) Pay(amount int) map[string]interface{} {
	res := p.paypal.SendPayment(amount)
	return map[string]interface{}{
		"id":     res["txnId"],
		"amount": res["total"],
	}
}

// // 3. Factory
var (
	registry = make(map[string]func() Adapter)
	mu       sync.RWMutex
)

func Register(name string, creator func() Adapter) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := registry[name]; exists {
		panic("Already registered: " + name)
	}
	registry[name] = creator
}

func Create(name string) Adapter {
	mu.RLock()
	creator, exists := registry[name]
	mu.RUnlock()
	if !exists {
		panic("unknown provider: " + name)
	}
	return creator()
}

//// 4. Builder

type Payment struct {
	ID     string
	Amount int
	Status string
}

type PaymentBuilder struct {
	p Payment
}

func (b *PaymentBuilder) SetID(id string) *PaymentBuilder {
	b.p.ID = id
	return b
}

func (b *PaymentBuilder) SetAmount(amount int) *PaymentBuilder {
	b.p.Amount = amount
	return b
}

func (b *PaymentBuilder) SetStatus(status string) *PaymentBuilder {
	b.p.Status = status
	return b
}

func (b *PaymentBuilder) Build() Payment {
	if b.p.ID == "" {
		panic("invalid payment")
	}
	return b.p
}

type Strategy interface {
	Process(p Payment) string
}

type SuccessStrategy struct{}

func (s SuccessStrategy) Process(p Payment) string {
	return "SUCCESS: " + p.ID
}

type FailureStrategy struct{}

func (f FailureStrategy) Process(p Payment) string {
	return "FAILED: " + p.ID
}

// Registry
var strategyRegistry = map[string]func() Strategy{}

func RegisterStrategy(name string, creator func() Strategy) {
	strategyRegistry[name] = creator
}

func GetStrategy(name string) Strategy {
	return strategyRegistry[name]()
}

//// 6. Main

func main() {
	Register("stripe", func() Adapter {
		return StripeAdapter{stripe: Stripe{}}
	})

	Register("paypal", func() Adapter {
		return PayPalAdapter{paypal: PayPal{}}
	})

	RegisterStrategy("SUCCESS", func() Strategy {
		return SuccessStrategy{}
	})

	RegisterStrategy("FAILED", func() Strategy {
		return FailureStrategy{}
	})

	adapter := Create("stripe")

	raw := adapter.Pay(1000)

	payment := (&PaymentBuilder{}).
		SetID(raw["id"].(string)).
		SetAmount(raw["amount"].(int)).
		SetStatus("SUCCESS").
		Build()

	strategy := GetStrategy(payment.Status)

	fmt.Println(strategy.Process(payment))
}

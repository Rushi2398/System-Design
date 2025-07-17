package main

type ConcreteObserver struct {
	Sub  Subject
	Name string
}

func (co *ConcreteObserver) Update(data string) {
	if co.Sub != nil {
		co.Sub.GetState(co.Name)
	}
}

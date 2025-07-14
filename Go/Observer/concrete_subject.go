package main

type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *ConcreteSubject) Unregister(o Observer) {
	for i, obs := range s.observers {
		if obs == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify(data string) {
	for _, o := range s.observers {
		o.Update(data)
	}
}

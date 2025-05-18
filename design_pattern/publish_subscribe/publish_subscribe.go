package publish_subscribe

import (
	"fmt"
	"github.com/samber/lo"
)

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

type Observer interface {
	Update()
}

type Boss struct {
	Action    string
	Observers []Observer
}

// Attach adds observer to Observers.
func (b *Boss) Attach(observer Observer) {
	b.Observers = append(b.Observers, observer)
}

// Detach remover observer from Observers.
func (b *Boss) Detach(observer Observer) {
	b.Observers = lo.Filter(b.Observers, func(item Observer, index int) bool {
		return item != observer
	})
}

// Notify make every observer call Update function.
func (b *Boss) Notify() {
	for _, obj := range b.Observers {
		obj.Update()
	}
}

// Make sure the object of Boss has implemented the Subject interface.
var _ Subject = &Boss{}

type Worker struct {
	Name string
}

func (w Worker) Update() {
	fmt.Println(w.Name, " is working")
}

var _ Observer = &Worker{}

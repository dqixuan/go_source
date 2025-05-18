package publish_subscribe

import "fmt"

/*
问题：生产者和消费者双向耦合，抽象生成者和消费者，实现解耦
*/
type Secretary struct {
	Action    string
	Observers []StockObserver
}

func (s *Secretary) Attach(observer StockObserver) {
	s.Observers = append(s.Observers, observer)
}

func (s *Secretary) Notify() {
	for _, observer := range s.Observers {
		observer.Update()
	}
}

type StockObserver struct {
	Name string
	Sub  Secretary
}

func NewStockObserver(name string, sub Secretary) *StockObserver {
	return &StockObserver{
		Name: name,
		Sub:  sub,
	}
}

func (s *StockObserver) Update() {
	fmt.Printf("{%s} {%s}关闭股票行情，继续工作\n", s.Sub.Action, s.Name)
}

func PublishSubscribeFirst() {
	s := Secretary{
		Action:    "老板回来了",
		Observers: nil,
	}
	stock1 := NewStockObserver("first name1", s)
	stock2 := NewStockObserver("first name2", s)
	s.Attach(*stock1)
	s.Attach(*stock2)
	s.Notify()
}

package main

import "fmt"

// 观察者观察目标
type Observer interface {
	observe()
}

type ObserverA struct {
	target *Target
}

type ObserverB struct {
	target *Target
}

type Target struct {
	state     int
	observers []Observer //目标被那些观察者观察
}

func (target *Target) notifyAll() {
	for i := 0; i < len(target.observers); i++ {
		target.observers[i].observe()
	}
}

// 修改状态后通知所有观察者
func (target *Target) modSate(state int) {
	target.state = state
	target.notifyAll()
}

func (target *Target) register(observer Observer) {
	target.observers = append(target.observers, observer)
}

func (oa *ObserverA) observe() {
	fmt.Println("observer A, target state is changed into ", oa.target.state)
}

func (ob *ObserverB) observe() {
	fmt.Println("observer B, target state is changed into ", ob.target.state)
}

func main() {
	t := &Target{state: 1, observers: []Observer{}}
	observerA := &ObserverA{t}
	observerB := &ObserverB{t}
	t.register(observerA)
	t.register(observerB)
	t.modSate(2)
}

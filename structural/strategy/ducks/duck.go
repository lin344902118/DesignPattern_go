package ducks

import "fmt"

type display interface {
	Display()
}

type Duck struct {
	flyBehavior FlyBehavior
	quackBehavior QuackBehavior
}

func (self *Duck) PerformFly() {
	self.flyBehavior.Fly()
}

func (self *Duck) PerformQuack() {
	self.quackBehavior.Quack()
}

func (self *Duck) Swim() {
	fmt.Println("All duck float, even decoys!")
}

func (self *Duck) SetFlyBehavior(behavior FlyBehavior) {
	self.flyBehavior = behavior
}

func (self *Duck) SetQuackBehavior(behavior QuackBehavior) {
	self.quackBehavior = behavior
}
package ducks

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct {
}

func (self *FlyWithWings) Fly() {
	fmt.Println("I'm flying")
}

type FlyNoWay struct {
}

func (self *FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}

type FlyRocketPowered struct {
}

func (self *FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket!")
}
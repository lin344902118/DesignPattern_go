package ducks

import "fmt"

type MallardDuck struct {
	Duck
}

func (self *MallardDuck) Display() {
	fmt.Println("I'm a real mallard duck")
}

func NewMallardDuck() MallardDuck {
	return MallardDuck{Duck{&FlyWithWings{}, &Quack{}}}
}
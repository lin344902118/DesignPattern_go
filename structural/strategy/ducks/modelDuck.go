package ducks

import "fmt"

type ModelDuck struct {
	Duck
}

func (self *ModelDuck) Display() {
	fmt.Println("I'm a model duck")
}

func NewModelDuck() ModelDuck {
	return ModelDuck{Duck{&FlyNoWay{}, &Quack{}}}
}
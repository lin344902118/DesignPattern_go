package ducks

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct {
}

func (self *Quack) Quack() {
	fmt.Println("Quack")
}

type MuteQuack struct {
}

func (self *MuteQuack) Quack() {
	fmt.Println("<<Silence>>")
}

type Squeak struct {
}

func (self *Squeak) Quack() {
	fmt.Println("Squeak")
}

package main

import "DesignPattern_go/structural/strategy/ducks"

func main() {
	mallardDuck := ducks.NewMallardDuck()
	mallardDuck.PerformQuack()
	mallardDuck.PerformFly()
	modelDuck := ducks.NewModelDuck()
	modelDuck.PerformFly()
	modelDuck.SetFlyBehavior(&ducks.FlyRocketPowered{})
	modelDuck.PerformFly()
}



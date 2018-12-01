package main

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyingBehavior struct {
	FlyBehavior
}

type NoFlyBehavior struct {
	FlyBehavior
}

func (d FlyingBehavior) Fly() {
	fmt.Println("Duck is flying")
}

func (d NoFlyBehavior) Fly() {
	fmt.Println("Duck can't fly")
}

type IDuck interface {
	Fly()
}

type Duck struct {
	fb FlyBehavior
}

func (d Duck) Fly() {
	d.fb.Fly()
}

func playWithDuck(d IDuck) {
	d.Fly()
	//d.Quack()
}

func main() {
	d1 := Duck{FlyingBehavior{}}
	d2 := Duck{NoFlyBehavior{}}

	playWithDuck(d1)
	playWithDuck(d2)
}

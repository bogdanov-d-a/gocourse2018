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

type QuackBehavior func()

type IDuck interface {
	Fly()
	Quack()
}

type Duck struct {
	fb FlyBehavior
	qb QuackBehavior
}

func (d Duck) Fly() {
	d.fb.Fly()
}

func (d Duck) Quack() {
	d.qb()
}

func playWithDuck(d IDuck) {
	d.Fly()
	d.Quack()
}

func main() {
	d1 := Duck{FlyingBehavior{}, func() { fmt.Println("Quack") }}
	d2 := Duck{NoFlyBehavior{}, func() { fmt.Println("Squeek") }}

	playWithDuck(d1)
	playWithDuck(d2)
}

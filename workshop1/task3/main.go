package main

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyingBehavior struct{}

type NoFlyBehavior struct{}

func (FlyingBehavior) Fly() {
	fmt.Println("Duck is flying")
}

func (NoFlyBehavior) Fly() {
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

func squeek() {
	fmt.Println("Squeek")
}

func main() {
	d1 := Duck{FlyingBehavior{}, func() { fmt.Println("Quack") }}
	d2 := Duck{NoFlyBehavior{}, squeek}

	playWithDuck(d1)
	playWithDuck(d2)
}

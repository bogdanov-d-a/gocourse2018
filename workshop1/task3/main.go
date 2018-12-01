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

type MallardDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{Duck{FlyingBehavior{}, func() { fmt.Println("Quack") }}}
}

type RubberDuck struct {
	Duck
}

func squeek() {
	fmt.Println("Squeek")
}

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{Duck{NoFlyBehavior{}, squeek}}
}

type WoodenDuck struct {
	Duck
}

func NewWoodenDuck() *WoodenDuck {
	return &WoodenDuck{Duck{NoFlyBehavior{}, func() { fmt.Println("Silent") }}}
}

func playWithDuck(d IDuck) {
	d.Fly()
	d.Quack()
}

func main() {
	mallardDuck := NewMallardDuck()
	rubberDuck := NewRubberDuck()
	woodenDuck := NewWoodenDuck()

	playWithDuck(mallardDuck)
	playWithDuck(rubberDuck)
	playWithDuck(woodenDuck)
}

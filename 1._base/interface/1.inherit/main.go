package main

import "fmt"

type Eater interface {
	Eat()
}

type Runner interface {
	Run()
}

type Animal interface {
	Eater
	Runner
}

type Dog struct {
	Name string
}

func (d *Dog) Eat() {
	fmt.Printf("%s is eating\n", d.Name)
}

func (d *Dog) Run() {
	fmt.Printf("%s is running\n", d.Name)
}

func ShowAnimalEat(animal Animal) {
	animal.Eat()
}

func ShowwAnimalRun(animal Animal) {
	animal.Run()
}

func ShowEaterEat(eater Eater) {
	eater.Eat()
}

func ShowRunnerRun(runner Runner) {
	runner.Run()
}

func main() {
	dog := Dog{Name: "Kenny"}
	ShowAnimalEat(&dog)
	ShowwAnimalRun(&dog)
	ShowEaterEat(&dog)
	ShowRunnerRun(&dog)
}

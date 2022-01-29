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
	fmt.Printf("DOG:%s is eating\n", d.Name)
}

func (d *Dog) Run() {
	fmt.Printf("DOG:%s is running\n", d.Name)
}

type Cat struct {
	Name string
}

func (c *Cat) Eat() {
	fmt.Printf("%s is eating\n", c.Name)
}

func (c *Cat) Run() {
	fmt.Printf("%s is running\n", c.Name)
}

func ShowEat(animal Animal) {
	animal.Eat()
}

func ShowRun(animal Animal) {
	animal.Run()
}

func ShowEat2(eater Eater) {
	eater.Eat()
}

func ShowRun2(runner Runner) {
	runner.Run()
}

func main() {
	animals := [...]Animal{
		&Dog{Name: "Kenny"},
		&Cat{Name: "Nicole"},
	}

	for _, animal := range animals {
		fmt.Println(animal)
	}
	for _, animal := range animals {
		switch animal.(type) {
		case *Dog:
			fmt.Println(animal.(*Dog).Name)
		case *Cat:
			fmt.Println(animal.(*Cat).Name)
		default:
			fmt.Println("you are not animal!!")
		}
	}

	
	fmt.Println("----------------")
	instances := [...]interface{}{
		123,
		"Hello World",
		&Dog{Name: "Kenny"},
		&Cat{Name: "Nicole"},
	}

	for _, instance := range instances {
		fmt.Println(instance)
	}
	for _, instance := range instances {
		switch instance.(type) {
		case *Dog:
			fmt.Println(instance.(*Dog).Name)
			instance.(*Dog).Run()
		case *Cat:
			fmt.Println(instance.(*Cat).Name)
		default:
			fmt.Println("you are not animal!!", instance)
		}
	}
}

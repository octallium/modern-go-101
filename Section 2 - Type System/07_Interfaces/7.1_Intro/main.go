package main

import "fmt"

type Animal interface {
	Walk()
	Talk()
}

type Dog struct {
	breed string
	name  string
}

type Cat struct {
	name string
	age  int
}

func (d *Dog) Walk() {
	fmt.Printf("%s is walking...\n", d.name)
}

func (d *Dog) Talk() {
	fmt.Printf("%s is saying Wuffff...\n", d.name)
}

func (c *Cat) Walk() {
	fmt.Printf("%s is walking...\n", c.name)
}

func (c *Cat) Talk() {
	fmt.Printf("%s is saying Meauuu...\n", c.name)
}

func greet(a Animal) {
	a.Talk()
}

func main() {
	var a1 Animal
	a1 = &Dog{breed: "Doberman", name: "Tommy"}
	a1.Talk()

	var a2 Animal
	a2 = &Cat{name: "Whisker", age: 3}
	a2.Walk()

	greet(a1)
	greet(a2)
}

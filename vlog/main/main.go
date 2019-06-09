package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int8
}

func main() {
	myPrint("Tom")
	person := Person{name: "Tim", age: 20}
	person.myPrint()
	p := Person{name: "Jerry", age: 30}
	person.myPrints(p)
	fmt.Printf("%v\n", person)
	fmt.Printf("%v\n", "vlog")
}

func myPrint(name string) {
	fmt.Printf("Hello %v\n", name)
}

func (person Person) myPrint() {
	fmt.Printf("%v\n", person)
}

func (person Person) myPrints(p Person) {
	fmt.Printf("%v\n", person)
	fmt.Printf("%v\n", p)
}

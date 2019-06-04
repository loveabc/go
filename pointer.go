package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int8
}

func main() {

	var person Person
	person.name = "Tom"
	person.age = 20
	fmt.Println(person)
	setPerson(&person)
	fmt.Println(person)

	p2 := Person{name: "Ketty", age: 20}
	fmt.Println(p2)

	var num int = 8
	fmt.Printf("num=%v\n", num)
	var ptr *int = &num
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("*ptr value=%v\n", *ptr)
	*ptr = 1000
	fmt.Printf("num=%v\n", num)
	num = 8

	if num > 10 {
		fmt.Println("1")
	} else if num < 5 {
		fmt.Println("2")
	} else {
		fmt.Println("3")
	}

	var v1, v2 int = max(1000, 2000)
	fmt.Println(v1, v2)

	numbers := [5]int{1, 2, 3, 4, 5}
	var age [5]int
	var i int
	for i = 0; i < 5; i++ {
		fmt.Println(numbers[i])
		age[i] = numbers[i] + 100
	}
	for i = 0; i < 5; i++ {
		fmt.Println(age[i])
	}
	names := []string{"a", "b"}
	for i = 0; i < 2; i++ {
		fmt.Println(names[i])
	}
}

func max(num1 int, num2 int) (int, int) {
	if num1 > num2 {
		return num1, num2
	}

	return num2, num1
}

func setPerson(person *Person) {
	(*person).name = "Tim"
}

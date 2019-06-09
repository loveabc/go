package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int8
}

func ExchangeString(s1, s2 string) {
	fmt.Printf("%v, %v\n", s1, s2)
	tmp := s1
	s1 = s2
	s2 = tmp
	fmt.Printf("%v, %v\n", s1, s2)
}

func ExchangeStringByPointer(s1, s2 *string) {
	fmt.Printf("%v, %v\n", *s1, *s2)
	tmp := *s1
	*s1 = *s2
	*s2 = tmp
	fmt.Printf("%v, %v\n", *s1, *s2)
}

func ExhangePerson(person Person) {
	fmt.Printf("%p\n", &person)
	fmt.Printf("%v\n", person)
	person.name = "Tom"
	fmt.Printf("%v\n", person)
}

func ExhangePersonByPointer(person *Person) {
	fmt.Printf("%v\n", *person)
	(*person).name = "Tom" //or person.name = "Tom"
	fmt.Printf("%v\n", *person)
}

func updateSlice(sli []string) []string {
	fmt.Printf("%p\n", &sli)
	return append(sli, "a")
}

func main() {
	s1 := "Tim"
	s2 := "Tom"
	ExchangeString(s1, s2)
	fmt.Printf("%v, %v\n", s1, s2)
	fmt.Println("=======")
	s1 = "Tim"
	s2 = "Tom"
	ExchangeStringByPointer(&s1, &s2)
	fmt.Printf("%v, %v\n", s1, s2)
	fmt.Printf("%v, %v\n", &s1, &s2)
	fmt.Printf("%v, %v\n", *&s1, *&s2)
	fmt.Println("====ttt===")
	person := Person{name: "Tim", age: 25}
	ExhangePerson(person)
	fmt.Printf("%p\n", &person)
	fmt.Println("====-----===")
	person = Person{name: "Tim", age: 25}
	ExhangePersonByPointer(&person)
	fmt.Printf("%v\n", person)
	fmt.Println("===========")
	// var sli []string = []string {"1", "2"}
	sli := []string{"1", "2"}
	fmt.Printf("%p\n", &sli)
	sli = updateSlice(sli)
	fmt.Println(sli)

	class := make(map[string]string)
	fmt.Printf("%p\n", &class)
	class["a"] = "A"
	fmt.Printf("%p\n", &class)
	fmt.Println(class)
}

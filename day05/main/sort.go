package main

import (
	"fmt"
	"sort"
	"time"
)

type Person struct{
	name string
	age int8
}

type sortByAge []Person
type sortByName []Person

func main() {
	i := 100
	j := 1000
	fmt.Println(i, j)
	swap(i, j)
	fmt.Println(i, j)

	names := []string{"10", "2", "3", "1", "5"}
	StringSort(names)

	ages := []int{1 ,2, 100, 20, 30, 50, 10}
	fmt.Println(ages)
	sort.Ints(ages)
	fmt.Println(ages)

	persons := []Person{
		{name: "Tom", age: 20},
		{name: "Jelly", age: 30},
		{name: "Tim", age: 15},
	}
	for _, p := range persons {
		fmt.Println(p)
	}
	fmt.Println("========")
	sort.Sort(sortByAge(persons))
	for _, p := range persons {
		fmt.Println(p)
	}
	fmt.Println("========")
	sort.Sort(sortByName(persons))
	for _, p := range persons {
		fmt.Println(p)
	}
}

func swap(i, j int) {
	fmt.Println(i, j)
	i, j = j, i
	fmt.Println(i, j)
}

func StringSort(names []string) {
	fmt.Println(names)
	sort.Strings(names)
	fmt.Println(names)

	du, _ := time.ParseDuration("3m20s")
	fmt.Printf("%T, %v\n", du, du)
}

func(person sortByAge) Len() int{
	return len(person)
}

func(person sortByAge) Less(i, j int) bool{
	return person[i].age < person[j].age
}

func(person sortByAge) Swap(i, j int) {
	person[i], person[j] = person[j], person[i]
}

func(person sortByName) Len() int{
	return len(person)
}

func(person sortByName) Less(i, j int) bool{
	return person[i].name < person[j].name
}

func(person sortByName) Swap(i, j int) {
	person[i], person[j] = person[j], person[i]
}
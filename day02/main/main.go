package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	var sl []int
	var arr2 = [...]int{1, 2, 3, 4, 5}
	var s2 = []int{1, 2, 3, 4, 5}
	fmt.Printf("%T, %d, %d\n", arr1, len(arr1), cap(arr1))
	fmt.Printf("%T, %d, %d\n", sl, len(sl), cap(sl))
	fmt.Printf("%T, %d, %d\n", arr2, len(arr2), cap(arr2))
	fmt.Printf("%T, %d, %d\n", s2, len(s2), cap(s2))
}
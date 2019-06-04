package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	s := []int{1, 2, 3, 4, 5, 6}
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	v1, v2 := <-ch, <-ch
	fmt.Println(v1, v2, v1+v2)
	fmt.Println(len(s))
	x, y := 100, 200
	nums := []int{1000, 2000, 3000}
	fmt.Println(x, y)
	fmt.Println(nums)
}

func sum(s []int, c chan int) {
	fmt.Println(s)
	var sum int = 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

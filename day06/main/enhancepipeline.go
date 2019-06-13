package main

import "fmt"

func main() {
	natural := make(chan int)
	square := make(chan int)
	go func () {
		for i := 0; i <= 100; i++ {
			natural <- i
		}
		defer close(natural)
	}()

	go func () {
		for x := range natural {
			square <- x * x
		}
		defer close(square)
	}()

	func () {
		for v := range square {
			fmt.Println(v)
		}
	}()
}

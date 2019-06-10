package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main...")
	go spinner(100 * time.Millisecond)
	v := fib(45)
	fmt.Printf("Fibonacci(%d) = %d\n", 45, v)
}
func spinner(delay time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x -1) + fib(x - 2)
}
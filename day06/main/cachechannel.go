package main

import (
	"fmt"
	"time"
)

//带缓存的channel
func main() {
	out := make(chan int, 5)
	go func () {
		for i := 0; i <= 10; i++ {
			out <- i
			fmt.Println("===")
		}
		close(out)
	}()

	func () {
		for v := range out {
			time.Sleep(2 * time.Second)
			fmt.Println(v)
		}
	}()
}

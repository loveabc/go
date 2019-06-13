package main

import (
	"fmt"
	"time"
)

func main() {
	natural := make(chan int)
	square := make(chan int)

	go func () {
		values := []int{0, 1, 2, 3, 4, 5}
		for _, v := range values {
			natural <- v
			time.Sleep(time.Second)
		}
		close(natural)
	}()

	go func () {
		for {
			v, ok := <- natural
			if !ok {
				break
			}
			square <- v * v
		}
		close(square)
	}()

	func () {
		for {
			v, ok := <- square
			if !ok {
				break
			}
			fmt.Println(v)
		}
	}()
}

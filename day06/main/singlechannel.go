package main

import "fmt"

//带方向的channel
func main() {
	natural := make(chan int)
	square := make(chan int)
	go counter(natural)
	go squarer(square, natural)
	printer(square)

}

func counter(out chan<- int) {
	for i := 0; i <= 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
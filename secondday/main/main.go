package main

import (
	"fmt"
)

type INT int

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))

	i := 10
	j := 20
	fmt.Println(i, j)
	i, j = j, i
	fmt.Println(i, j)

	g := gcd(100, 50)
	fmt.Println("gcd=", g)
	var m int = 100
	var n INT = 10
	var sum INT = INT(m) + n
	fmt.Println("sum=", sum)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func init() {
	fmt.Println("==123==")
}
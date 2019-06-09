package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, World"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Printf("%c, %c\n", s[0], s[7])
	str := s[:]
	fmt.Println(str)
	sliceStr := s[2:8]
	fmt.Println(sliceStr)
	fmt.Println(s == str)
	fmt.Println(&s)
	fmt.Println(&str)

	s = "你好，中国"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	iter()
}

func iter() {
	s := "Hello, World"
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}
	s = "你好，中国"
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}
}
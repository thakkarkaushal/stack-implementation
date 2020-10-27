package main

import (
	"fmt"

	"github.com/thakkarkaushal/stack-implementation/stack"
)

func main() {
	for i := 0; i < 2; i++ {
		s := stack.Push(i)
		fmt.Println(s)
	}

	s := stack.Pop()
	fmt.Println(s)
	s = stack.Pop()
	fmt.Println(s)

	s = stack.Pop()
	fmt.Println(s)

	s = stack.Pop()
	fmt.Println(s)
}

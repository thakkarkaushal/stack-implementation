package stack

import (
	"fmt"
)

// Stack store the int values.
var Stack []int

// Max limit to store the values.
var Max = 10

// Push function to add value to Stack array
func Push(value int) []int {
	stackLength := len(Stack)
	if stackLength == Max {
		fmt.Print("Stack is full.\n")
		return Stack
	}
	Stack = append(Stack, value)
	return Stack
}

// Pop remove the top element from the Stack array
func Pop() []int{
	stackLength := len(Stack)
	if stackLength == 0 {
		fmt.Println("Stack is empty cannot remove element from stack.")
		return Stack
	}
	fmt.Printf("%d removed from stack.\n", Stack[stackLength-1])
	Stack = Stack[:stackLength-1]
	return Stack
}

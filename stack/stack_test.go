package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func addToStack(endpoint int) []int{
	var o []int
	for i:=0; i < endpoint ; i++ {
		o = Push(i)
	}
	return o
}

func TestPush(t *testing.T) {
	var expectedOutput = []int{50}
	output := Push(50)
	fmt.Println(output)
	if !reflect.DeepEqual(expectedOutput, output){
		t.Fail()
	}
}

func TestPop(t *testing.T){
	var expectedOutput = []int{0,1}
	addToStack(3)
	output := Pop()
	fmt.Println(output)
	if !reflect.DeepEqual(expectedOutput, output){
		t.Fail()
	}
}

func TestEmptyPop(t *testing.T){
	var expectedOutput = []int{}
	output := Pop()
	fmt.Println(output)
	if reflect.DeepEqual(expectedOutput, output){
		t.Fail()
	}
}

func TestMaxPush(t *testing.T){
	var expectedOutput = []int{0,1,2,3,4,5,6,7,8,9}
	output := addToStack(11)
	fmt.Println(output)
	if !reflect.DeepEqual(expectedOutput, output){
		t.Fail()
	}
}

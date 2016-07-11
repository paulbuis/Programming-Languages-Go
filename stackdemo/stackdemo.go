package main

import (
	"fmt"
	"github.com/paulbuis/stack"
)

func main() {
	fmt.Println("trying a SliceStack")
	st0 := stack.NewSliceStack()
	useStack(&st0)

	fmt.Println("trying a ListStack")
	st1 := stack.NewListStack()
	useStack(&st1)

}

func useStack(s stack.Stack) {
	s.Push(1)
	s.Push(2)
	val, err := s.Peek()
	if err == nil {
		fmt.Println(val)
	} else {
		return
	}

	val, err = s.Pop()
	if err == nil {
		fmt.Println(val)
	} else {
		return
	}
	val, err = s.Pop()
	if err == nil {
		fmt.Println(val)
	}
}

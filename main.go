package main

import (
	"fmt"
	"github.com/Javlopez/dsa/pkg/stack"
)

func main() {
	stack1 := stack.New[int]()
	// stack1 := stack.Stack[int]{}
	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)
	fmt.Println(stack1)
	poped, err := stack1.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(poped)
	fmt.Println(stack1)
	fmt.Println(stack1.Pop()) // 3

	poped, err = stack1.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stack1)
	fmt.Println(stack1.Pop()) // 2
	poped, err = stack1.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stack1)
	fmt.Println(stack1.Pop()) // 1
}

package main

import "fmt"

type stack []string

func new(v ...string) stack {
	return v
}

func (s *stack) pop() string {
	result := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return result
}

func (s *stack) append(v string) {
	*s = append(*s, v)
}

func test_lifo() {
	stack := new("a", "b", "c")
	fmt.Println(stack)
	stack.append("d")
	fmt.Println(stack)
	r := stack.pop()
	fmt.Println(r)
	fmt.Println(stack)
}

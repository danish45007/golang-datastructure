package main

import "fmt"

/*
	LIFO(Last In First Out)
*/

// Stack is a struct that holds slice
type Stack struct {
	array []int
}

func (s *Stack) Push(item int) {
	s.array = append(s.array, item)
}

func (s *Stack) Pop() {
	l := len(s.array) - 1
	fmt.Println("POP OUT", s.array[l])
	s.array = s.array[:l]
}

func main() {
	stack := Stack{}
	data := []int{1, 2, 3, 5}
	for _, d := range data {
		stack.Push(d)
		fmt.Println(stack)
	}
	for i := 0; i < 4; i++ {
		stack.Pop()
		fmt.Println(stack)
	}

}

package stacks

import (
	"errors"
	"fmt"
)

type stack struct {
	items []int
}

func (s *stack) push(newItem int) {
	s.items = append(s.items, newItem)
}

func (s *stack) pop() (int, error) {
	if len(s.items) > 0 {
		removed := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return removed, nil
	} else {
		return 0, errors.New("empty stack")
	}
}

func (s stack) length() int {
	return len(s.items)
}

func InitStackExample() {
	myStack := stack{}
	myStack.push(10)
	myStack.push(20)
	myStack.push(30)
	fmt.Println(myStack)
	fmt.Println("Length:", myStack.length())
	removed1, _ := myStack.pop()
	fmt.Println("Remoevd:", removed1)
	fmt.Println(myStack)
	myStack.pop()
	myStack.pop()
	removedErr, err := myStack.pop()
	fmt.Println(removedErr, err)
	fmt.Println(myStack)
}

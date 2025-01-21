package stack

import "fmt"

type Stack []int

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() int {
	if len((*s)) == 0 {
		panic("stack is empty")
	}
	popElement := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popElement
}
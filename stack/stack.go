package stack

type Stack[T any] []T

func (s *Stack[T]) Push(data T) {
	*s = append(*s, data)
}

func (s *Stack[T]) Pop() T {
	if len((*s)) == 0 {
		panic("stack is empty")
	}
	popElement := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popElement
}

func (s *Stack[T]) Peek() T {
	if len((*s)) == 0 {
		panic("stack is empty")
	}
	return (*s)[len((*s))-1]
}

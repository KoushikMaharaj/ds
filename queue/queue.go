package queue

type Queue[T any] []T

func (q *Queue[T]) Enqueue(data T) {
	*q = append(*q, data)
}

func (q *Queue[T]) Dequeue() T {
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}

func (q *Queue[T]) IsEmpty() bool {
	return len((*q)) == 0
}

package marsType

import "github.com/marsli9945/mars-go/marsLog"

type Queue[T any] []T

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	if q.IsEmpty() {
		marsLog.Logger().Error("queue is empty")
		return *new(T)
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

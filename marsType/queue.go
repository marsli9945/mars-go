package marsType

type Queue[T any] []T

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	if q.IsEmpty() {
		return *new(T)
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

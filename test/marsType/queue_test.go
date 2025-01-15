package marsType

import (
	"github.com/marsli9945/mars-go/marsType"
	"reflect"
	"testing"
)

func TestQueue_Push(t *testing.T) {
	q := marsType.Queue[int]{}
	q.Push(1)
	q.Push(2)
	q.Push(3)

	expected := marsType.Queue[int]{1, 2, 3}
	if !reflect.DeepEqual(q, expected) {
		t.Errorf("Push: expected %v, got %v", expected, q)
	}
}

func TestQueue_Pop(t *testing.T) {
	q := marsType.Queue[int]{1, 2, 3}
	popped := q.Pop()

	if popped != 1 {
		t.Errorf("Pop: expected 1, got %d", popped)
	}

	expected := marsType.Queue[int]{2, 3}
	if !reflect.DeepEqual(q, expected) {
		t.Errorf("Pop: expected %v, got %v", expected, q)
	}

	// 测试从空队列中弹出
	q = marsType.Queue[int]{}
	popped = q.Pop()
	if popped != 0 {
		t.Errorf("Pop from empty queue: expected 0, got %d", popped)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := marsType.Queue[int]{}
	if !q.IsEmpty() {
		t.Errorf("IsEmpty: expected true, got false")
	}

	q.Push(1)
	if q.IsEmpty() {
		t.Errorf("IsEmpty: expected false, got true")
	}
}

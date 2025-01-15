package marsType

import "sync"

type Set[T string | int] struct {
	mu    sync.Mutex
	items map[T]struct{}
}

func NewSet[T string | int]() Set[T] {
	return Set[T]{items: make(map[T]struct{})}
}

func (s *Set[T]) Add(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[item] = struct{}{}
}

func (s *Set[T]) AddAll(items []T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, item := range items {
		s.items[item] = struct{}{}
	}
}

func (s *Set[T]) Remove(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, item)
}

func (s *Set[T]) Contains(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.items[item]
	return exists
}

func (s *Set[T]) ToList() Array[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	list := make(Array[T], 0, len(s.items))
	for k := range s.items {
		list = append(list, k)
	}
	return list
}

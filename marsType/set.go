package marsType

import "sync"

type Set[T string | int] struct {
	mu    sync.Mutex
	items map[T]struct{}
}

func NewSet[T string | int](items ...T) *Set[T] {
	set := Set[T]{items: make(map[T]struct{})}
	set.Add(items...)
	return &set
}

func (s *Set[T]) Add(items ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, item := range items {
		s.items[item] = struct{}{}
	}
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

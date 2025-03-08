package goset

import "fmt"

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

func (s Set[T]) Contains(elem T) bool {
	_, ok := s[elem]
	return ok
}

func (s Set[T]) Get() T {
	for elem := range s {
		return elem
	}
	var t T
	return t
}

func (s Set[T]) TryGet() (T, error) {
	for elem := range s {
		return elem, nil
	}
	var t T
	return t, fmt.Errorf("set is empty")
}

func (s Set[T]) TryAdd(elem T) bool {
	if _, ok := s[elem]; ok {
		return false
	}
	s[elem] = struct{}{}
	return true
}

func (s Set[T]) TryRemove(elem T) bool {
	if _, ok := s[elem]; ok {
		delete(s, elem)
		return true
	}
	return false
}

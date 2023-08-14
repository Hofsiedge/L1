package main

import (
	"fmt"
	"strings"
)

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

// Set[T] represents a set of T as a map with T keys and empty struct values
//
// struct{}{} is 0 byte and maps already implement
// the set functionality over keys
type Set[T comparable] struct {
	values map[T]struct{}
}

// String() to implement Stringer
func (s *Set[T]) String() string {
	var zero T
	values := make([]string, 0, len(s.values))
	for x := range s.values {
		values = append(values, fmt.Sprintf("%v", x))
	}
	return fmt.Sprintf("Set[%T]{%s}", zero, strings.Join(values, ", "))
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(obj T) {
	s.values[obj] = struct{}{}
}

func (s *Set[T]) Remove(obj T) {
	delete(s.values, obj)
}

func (s *Set[T]) Contains(obj T) bool {
	_, ok := s.values[obj]
	return ok
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	var a, b *Set[T]
	// choosing a to be the smallest map to reduce the number of iterations
	if len(s.values) <= len(other.values) {
		a, b = s, other
	} else {
		a, b = other, s
	}
	result := NewSet[T]()
	for obj := range a.values {
		if b.Contains(obj) {
			result.Add(obj)
		}
	}
	return result
}

func (s *Set[T]) Union(other ...*Set[T]) *Set[T] {
	sets := make([]*Set[T], 0, len(other)+1)
	sets = append(sets, s)
	sets = append(sets, other...)
	result := NewSet[T]()
	for _, m := range sets {
		for obj := range m.values {
			result.Add(obj)
		}
	}
	return result
}

func main() {
	A := NewSet[int]()
	for _, x := range []int{10, 2, 10, 1, 3, -5} {
		A.Add(x)
	}
	B := NewSet[int]()
	for _, x := range []int{10, 8, 1, 4, 4} {
		B.Add(x)
	}
	C := A.Intersect(B)
	fmt.Printf("%v & %v = %v", A, B, C)
}

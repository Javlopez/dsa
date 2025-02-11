package stack

import "fmt"

type Stack[T any] struct {
	elements []T
}

var ErrStackEmpty = fmt.Errorf("stack is empty")

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, ErrStackEmpty
	}

	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return &top, nil
}

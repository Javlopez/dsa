package linkedList

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
}

type Equals[T any] func(a, b T) bool

var ErrNodeNotFound = fmt.Errorf("node not found")

func (l *LinkedList[T]) Append(element T) {
	node := &Node[T]{Value: element}
	if l.Head == nil {
		l.Head = node
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
}

func (l *LinkedList[T]) Search(element T, equals Equals[T]) (*Node[T], error) {
	if l.Head == nil {
		return nil, ErrNodeNotFound
	}

	for current := l.Head; current != nil; current = current.Next {
		if equals(current.Value, element) {
			return current, nil
		}

		if current.Next == nil {
			break
		}
		current = current.Next
	}

	return nil, ErrNodeNotFound
}

func (l *LinkedList[T]) RemoveElement(element T, equals Equals[T]) error {
	if l.Head == nil {
		return ErrNodeNotFound
	}

	if equals(l.Head.Value, element) {
		l.Head = l.Head.Next
		return nil
	}

	current := l.Head
	for current.Next != nil {
		if equals(current.Next.Value, element) {
			current.Next = current.Next.Next
			if current.Next == nil {
				current.Next = nil
			}
			return nil
		}
		current = current.Next
	}
	return nil
}

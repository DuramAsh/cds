package stack

import (
	"github.com/duramash/cds/node"
)

type Stack[T comparable] struct {
	top *node.Node[T]
}

func New[T comparable]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Size() (size int) {
	for n := s.top; n != nil; n = n.Next {
		size++
	}

	return
}

func (s *Stack[T]) Pop() (value T) {
	if s.top == nil {
		panic("Stack is empty")
	}

	value = s.top.Value
	s.top = s.top.Next

	return
}

func (s *Stack[T]) Peek() (value T) {
	if s.top == nil {
		panic("Stack is empty")
	}

	value = s.top.Value

	return
}

func (s *Stack[T]) Empty() bool {
	return s.top == nil
}

func (s *Stack[T]) Push(value T) {
	s.top = node.New(value, s.top, nil)
}

func (s *Stack[T]) Clear() {
	s.top = nil
}

package queue

import (
	"github.com/duramash/cds/node"
)

type Queue[T comparable] struct {
	head *node.Node[T]
	tail *node.Node[T]
}

func New[T comparable]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Size() (size int) {
	for n := q.head; n != nil; n = n.Next {
		size++
	}

	return
}

func (q *Queue[T]) Enqueue(value T) {
	if q.head == nil {
		q.head = node.New(value, nil, nil)
		q.tail = q.head
		return
	}

	q.tail.Next = node.New(value, nil, nil)
	q.tail = q.tail.Next
}

func (q *Queue[T]) Dequeue() (value T) {
	if q.head == nil {
		panic("Queue is empty")
	}

	value = q.head.Value
	q.head = q.head.Next

	return
}

func (q *Queue[T]) Peek() (value T) {
	if q.head == nil {
		panic("Queue is empty")
	}

	value = q.head.Value

	return
}

func (q *Queue[T]) Empty() bool {
	return q.head == nil
}

func (q *Queue[T]) Clear() {
	q.head = nil
	q.tail = nil
}

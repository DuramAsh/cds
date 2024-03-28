package deque

import (
	"github.com/duramash/cds/node"
)

type Deque[T comparable] struct {
	head *node.Node[T]
	tail *node.Node[T]
}

func New[T comparable]() Deque[T] {
	return Deque[T]{}
}

func (d *Deque[T]) Size() (size int) {
	for n := d.head; n != nil; n = n.Next {
		size++
	}

	return
}

func (d *Deque[T]) PushFront(value T) {
	if d.head == nil {
		d.head = node.New(value, nil, nil)
		d.tail = d.head
		return
	}

	d.head.Prev = node.New(value, d.head, nil)
	d.head = d.head.Prev
}

func (d *Deque[T]) PushBack(value T) {
	if d.head == nil {
		d.head = node.New(value, nil, nil)
		d.tail = d.head
		return
	}

	d.tail.Next = node.New(value, nil, d.tail)
	d.tail = d.tail.Next
}

func (d *Deque[T]) PopFront() (value T) {
	if d.head == nil {
		panic("Deque is empty")
	}

	value = d.head.Value
	d.head = d.head.Next

	if d.head == nil {
		d.tail = nil
		return
	}

	d.head.Prev = nil

	return
}

func (d *Deque[T]) PopBack() (value T) {
	if d.head == nil {
		panic("Deque is empty")
	}

	value = d.tail.Value
	d.tail = d.tail.Prev

	if d.tail == nil {
		d.head = nil
		return
	}

	d.tail.Next = nil

	return
}

func (d *Deque[T]) PeekFront() (value T) {
	if d.head == nil {
		panic("Deque is empty")
	}

	value = d.head.Value

	return
}

func (d *Deque[T]) PeekBack() (value T) {
	if d.head == nil {
		panic("Deque is empty")
	}

	value = d.tail.Value

	return
}

func (d *Deque[T]) Empty() bool {
	return d.head == nil
}

func (d *Deque[T]) Clear() {
	d.head = nil
	d.tail = nil
}

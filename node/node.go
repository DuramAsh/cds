package node

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

func New[T comparable](value T, next, prev *Node[T]) *Node[T] {
	return &Node[T]{
		Value: value,
		Next:  next,
		Prev:  prev,
	}
}

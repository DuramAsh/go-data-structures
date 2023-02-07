package stack

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func MakeNode[T any](val T) *Node[T] {
	return &Node[T]{Val: val, Next: nil}
}

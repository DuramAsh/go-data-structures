package singlylinkedlist

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func makeNode[T any](val T) *Node[T] {
	return &Node[T]{Val: val, Next: nil}
}

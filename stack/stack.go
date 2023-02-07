package stack

import "fmt"

type Stack[T comparable] struct {
	top  *Node[T]
	size int
}

func MakeStack[T comparable]() Stack[T] {
	return Stack[T]{top: nil, size: 0}
}

func (s *Stack[T]) setTop(top *Node[T]) {
	s.top = top
}

func (s *Stack[T]) setSize(size int) {
	s.size = size
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack[T]) Push(val T) {
	newNode := MakeNode(val)
	if !s.IsEmpty() {
		newNode.Next = s.top
	}
	s.setTop(newNode)
	s.setSize(s.Size() + 1)
}

func (s *Stack[T]) Peek() *Node[T] {
	return s.top
}

func (s *Stack[T]) Pop() (*Node[T], error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	toRet := s.Peek()
	s.setSize(s.Size() - 1)
	s.setTop(toRet.Next)
	return toRet, nil
}

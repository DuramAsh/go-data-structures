package sll

import (
	"fmt"
)

//type LinkedList[T comparable] interface {
//	PushHead()
//	PushTail()
//	Reverse()
//	Size()
//	Head()
//}

type SLL[T comparable] struct {
	head *Node[T]
}

func MakeSLL[T comparable]() SLL[T] {
	return SLL[T]{head: nil}
}

func (sll *SLL[T]) GetHead() *Node[T] {
	return sll.head
}

func (sll *SLL[T]) SetHead(node *Node[T]) {
	sll.head = node
}

func (sll *SLL[T]) PushHead(val T) {
	newNode := MakeNode(val)
	if sll.GetHead() != nil {
		newNode.Next = sll.GetHead()
		sll.SetHead(newNode)
	} else {
		sll.SetHead(newNode)
	}
}

func (sll *SLL[T]) PushTail(val T) {
	newNode := MakeNode(val)
	if sll.GetHead() != nil {
		cur := sll.GetHead()
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newNode
	} else {
		sll.SetHead(newNode)
	}
}

func (sll *SLL[T]) Print() {
	cur := sll.GetHead()
	fmt.Print("[")
	for cur != nil {
		fmt.Print(cur.Val)
		if cur.Next != nil {
			fmt.Print(",")
		}
		cur = cur.Next
	}
	fmt.Print("]\n")
}

func (sll *SLL[T]) Reverse() {
	var prev *Node[T]
	cur := sll.GetHead()
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	sll.SetHead(prev)
}

func (sll *SLL[T]) Size() int {
	cur, cnt := sll.GetHead(), 0
	for cur != nil {
		cur = cur.Next
		cnt++
	}
	return cnt
}

func (sll *SLL[T]) Insert(val T, pos int) error {
	size := sll.Size()
	if pos < 0 || pos > size {
		return fmt.Errorf("position index out of range")
	}
	newNode := MakeNode(val)
	if pos == 0 {
		sll.PushHead(val)
	} else if pos == size {
		sll.PushTail(val)
	} else {
		cur := sll.GetHead()
		for pos > 1 {
			pos--
			cur = cur.Next
		}
		newNode.Next = cur.Next
		cur.Next = newNode
	}
	return nil
}

func (sll *SLL[T]) PopHead() (*Node[T], error) {
	if sll.Size() == 0 {
		return nil, fmt.Errorf("list is empty")
	}
	head := sll.GetHead()
	sll.SetHead(head.Next)
	return head, nil
}

func (sll *SLL[T]) PopTail() (*Node[T], error) {
	if sll.Size() == 0 {
		return nil, fmt.Errorf("list is empty")
	}
	tail := sll.GetHead()
	if tail.Next == nil {
		sll.SetHead(nil)
		return tail, nil
	}
	for tail.Next.Next != nil {
		tail = tail.Next
	}
	ret := tail.Next
	tail.Next = nil
	return ret, nil
}

// TODO: delete from index

func (sll *SLL[T]) Find(val T) *Node[T] {
	cur := sll.GetHead()
	for cur != nil {
		if cur.Val == val {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func (sll *SLL[T]) Sort(comp func(a, b T) bool) {
	cur := sll.GetHead()
	var ind *Node[T] = nil
	var temp T
	if sll.GetHead() == nil {
		return
	}
	for cur != nil {
		ind = cur.Next
		for ind != nil {
			if comp(ind.Val, cur.Val) {
				temp = cur.Val
				cur.Val = ind.Val
				ind.Val = temp
			}
			ind = ind.Next
		}
		cur = cur.Next
	}
}

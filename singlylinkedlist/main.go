package singlylinkedlist

import "fmt"

//type LinkedList[T any] interface {
//	PushHead()
//	PushTail()
//	Reverse()
//	Size()
//	Head()
//}

type SinglyLinkedList[T any] struct {
	head *Node[T]
}

func (sll *SinglyLinkedList[T]) getHead() *Node[T] {
	return sll.head
}

func (sll *SinglyLinkedList[T]) setHead(node *Node[T]) {
	sll.head = node
}

func (sll *SinglyLinkedList[T]) PushHead(val T) {
	newNode := makeNode(val)
	if sll.getHead() != nil {
		newNode.Next = sll.getHead()
		sll.setHead(newNode)
	} else {
		sll.setHead(newNode)
	}
}

func (sll *SinglyLinkedList[T]) PushTail(val T) {
	newNode := makeNode(val)
	if sll.getHead() != nil {
		cur := sll.getHead()
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newNode
	} else {
		sll.setHead(newNode)
	}
}

func (sll *SinglyLinkedList[T]) Print() {
	cur := sll.getHead()
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

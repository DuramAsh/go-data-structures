package main

import (
	"github.com/duramash/go-data-structures/singlylinkedlist/src/sll"
	"strconv"
)

func main() {
	l := sll.SLL[string]{}
	for i := 0; i < 10; i++ {
		l.PushTail(strconv.Itoa(i) + "))")
	}
	l.Print()
	l.Reverse()
	l.Sort(func(a string, b string) bool {
		return a < b
	})
	l.Print()
}

package main

import "fmt"

type list[k comparable] struct {
	head   *node[k]
	tail   *node[k]
	length int
}

func (l *list[k]) add(data k) {
	n := &node[k]{data: data}
	l.tail.next = n
	l.tail = n
}

func (l *list[k]) print() {
	for n := l.head; n != nil; n = n.next {
		fmt.Print(n.data)
	}
}

type node[k comparable] struct {
	data k
	next *node[k]
}

func newList[k comparable](data k) *list[k] {
	head := &node[k]{data: data}
	tail := head
	lst := &list[k]{head: head, tail: tail}
	return lst
}

func main() {
	lst := newList("hola")
	lst.add("mundo")

	lst.print()
}

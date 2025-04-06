package ds

import "cmp"

/*
*
https://leetcode.com/problems/merge-two-sorted-lists/description/
*/
type Node[k cmp.Ordered] struct {
	data k
	next *Node[k]
}

func NewNode[K cmp.Ordered](k K) *Node[K] {
	return &Node[K]{data: k}
}

type LinkedList[k cmp.Ordered] struct {
	head   *Node[k]
	tail   *Node[k]
	length int
}

func (lst *LinkedList[K]) Walker(TraverseFunc func(k K)) {
	for n := lst.head; n != nil; n = n.next {
		TraverseFunc(n.data)
	}
}

func (l *LinkedList[k]) Append(data k) *LinkedList[k] {
	n := &Node[k]{data: data}
	if l.head == nil {
		l.head = n
		l.tail = n
	}
	l.tail.next = n
	l.tail = n
	l.length++
	return l
}

/*
Merge You are given two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/
func (l *LinkedList[k]) Merge(l2 *LinkedList[k]) *LinkedList[k] {
	n1 := l.head
	n2 := l2.head
	merged := &LinkedList[k]{}
	for n1 != nil || n2 != nil {
		if n1 == nil {
			merged.Append(n2.data)
			n2 = n2.next
			continue
		}
		if n2 == nil {
			merged.Append(n1.data)
			n1 = n1.next
			continue
		}
		if n1.data > n2.data {
			merged.Append(n2.data)
			n2 = n2.next
			continue
		}
		merged.Append(n1.data)
		n1 = n1.next
	}
	return merged
}

func (l *LinkedList[k]) Copy() *LinkedList[k] {
	var newList *LinkedList[k]
	l.Walker(func(node k) {
		if newList != nil {
			newList.Append(node)
		} else {
			newList = NewLinkedList(node)
		}
	})
	return newList
}

func (l *LinkedList[k]) CopyV2() *LinkedList[k] {
	newLst := &LinkedList[k]{}
	if l.head == nil {
		return newLst
	}
	for n := l.head; n != nil; n = n.next {
		newNode := NewNode(n.data)
		if newLst.tail != nil {
			// anything that is currently pointing at the end (tail)
			// is going to point to the new node
			newLst.tail.next = newNode
			// update the current tail pointer to the new node
			newLst.tail = newNode
		} else {
			newLst.tail = newNode
			newLst.head = newNode
		}
		newLst.length++
	}
	return newLst
}

func NewLinkedList[K cmp.Ordered](data K) *LinkedList[K] {
	head := &Node[K]{data: data}
	tail := head
	lst := &LinkedList[K]{head: head, tail: tail}
	return lst
}

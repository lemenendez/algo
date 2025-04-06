package ds

type LeetListNode struct {
	Val  int
	Next *LeetListNode
}

func LeetMergeTwoLists(list1 *LeetListNode, list2 *LeetListNode) *LeetListNode {
	// manage trivial cases
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	n1 := list1
	n2 := list2

	var head *LeetListNode
	var tail *LeetListNode

	for n1 != nil || n2 != nil {
		newNode := &LeetListNode{}
		if head == nil {
			head = newNode
			tail = head
		} else {
			tail.Next = newNode
			tail = newNode
		}

		if n1 == nil {
			newNode.Val = n2.Val
			n2 = n2.Next
			continue
		}

		if n2 == nil {
			newNode.Val = n1.Val
			n1 = n1.Next
			continue
		}

		if n1.Val > n2.Val {
			newNode.Val = n2.Val
			n2 = n2.Next
			continue
		}

		newNode.Val = n1.Val
		n1 = n1.Next

	}

	return head
}

func ToLeet(a []int) *LeetListNode {
	var head *LeetListNode
	node := head
	for i := 0; i < len(a); i++ {
		if head == nil {
			head = &LeetListNode{Val: a[i]}
			node = head
			continue
		}
		node.Next = &LeetListNode{Val: a[i]}
		node = node.Next
	}

	return head
}

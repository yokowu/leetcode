package linked

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func Append(head *ListNode, val int) *ListNode {
	n := &ListNode{Val: val}
	if head == nil {
		head = n
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = n
	return head
}

func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for tmp := dummy; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummy.Next
}

func Show(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

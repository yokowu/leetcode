package linked

import (
	"fmt"
)

type Node struct {
	Next *Node
	Val  int
}

func New(val int) *Node {
	return &Node{Val: val}
}

func Add(head *Node, val int) *Node {
	newHead := &Node{Val: val, Next: head}
	return newHead
}

func Reverse(head *Node) *Node {
	var pre, next *Node
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// MidUp 输入链表头节点, 奇数长度返回中点, 偶数长度返回上中点
func MidUp(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	slow := head.Next
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// MidDown 输入链表头节点, 奇数长度返回中点, 偶数长度返回下中点
func MidDown(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head.Next
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func Show(head *Node) {
	fmt.Println("----------")
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

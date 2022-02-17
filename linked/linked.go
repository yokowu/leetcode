package linked

import "fmt"

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

func Show(head *Node) {
	fmt.Println("----------")
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

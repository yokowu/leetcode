package linked

import "fmt"

type Linked struct {
	len  int
	head *node
}

type node struct {
	next *node
	Val  int
}

func New() *Linked {
	return &Linked{
		len:  0,
		head: nil,
	}
}

func (l *Linked) Insert(val int) {
	n := &node{Val: val}
	n.next = l.head
	l.head = n
	l.len++
}

func (l *Linked) Append(val int) {
	n := &node{Val: val}
	if l.len == 0 {
		l.head = n
		l.len++
		return
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = n
	l.len++
}

func (l *Linked) Reverse() {
	var n *node
	for l.head != nil {
		next := l.head.next
		l.head.next = n
		n = l.head
		l.head = next
	}
	l.head = n
}

func (l *Linked) Show() {
	cur := l.head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.next
	}
}

func (l *Linked) Len() {
	fmt.Printf("Linked len: %d", l.len)
}

// // MidUp 输入链表头节点, 奇数长度返回中点, 偶数长度返回上中点
// func MidUp(head *Node) *Node {
// 	if head == nil || head.Next == nil || head.Next.Next == nil {
// 		return head
// 	}

// 	slow := head.Next
// 	fast := head.Next.Next
// 	for fast.Next != nil && fast.Next.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
// 	return slow
// }

// // MidDown 输入链表头节点, 奇数长度返回中点, 偶数长度返回下中点
// func MidDown(head *Node) *Node {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	slow := head.Next
// 	fast := head.Next
// 	for fast.Next != nil && fast.Next.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
// 	return slow
// }

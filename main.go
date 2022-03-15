package main

import (
	"fmt"

	"github.com/yokowu/leetcode/linked"
)

func main() {
	n := linked.ListNode{
		Val: 10,
	}
	n.Next = &linked.ListNode{
		Val: 5,
		Next: &linked.ListNode{
			Val:  3,
			Next: &n,
		},
	}

	b := linked.HasCycle(&n)
	fmt.Println(b)
}

package main

import (
	"github.com/yokowu/leetcode/linked"
)

func main() {
	head := linked.New(1)
	head = linked.Add(head, 2)
	head = linked.Add(head, 3)
	head = linked.Add(head, 4)
	head = linked.Add(head, 5)
	linked.Show(head)
	head = linked.Reverse(head)
	linked.Show(head)
}

package main

import (
	"github.com/yokowu/leetcode/linked"
)

func main() {
	l1 := linked.New(1)
	l1 = linked.Append(l1, 2)
	n := linked.New(3)
	l1 = linked.AppendNode(l1, n)
	l1 = linked.Append(l1, 6)

	linked.DeleteNode(n)
	linked.Show(l1)
}

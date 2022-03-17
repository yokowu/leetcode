package main

import (
	"fmt"

	"github.com/yokowu/leetcode/linked"
)

func main() {
<<<<<<< HEAD
	node := linked.New(1)
	node = linked.Append(node, 2)
	node = linked.Append(node, 3)
	node = linked.Append(node, 2)
	node = linked.Append(node, 1)

	b := linked.IsPalindrome2(node)
	fmt.Println(b)
=======
	l1 := linked.New(1)
	l1 = linked.Append(l1, 2)
	l1 = linked.Append(l1, 3)
	l1 = linked.Append(l1, 6)

	l2 := linked.New(2)
	l2 = linked.Append(l2, 3)
	l2 = linked.Append(l2, 5)

	n := linked.MergeTwoLists(l1, l2)
	linked.Show(n)
>>>>>>> linked merge
}

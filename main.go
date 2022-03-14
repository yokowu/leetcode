package main

import (
	"fmt"

	"github.com/yokowu/leetcode/linked"
)

func main() {
	root := linked.New(6)
	root = linked.Add(root, 5)
	root = linked.Add(root, 4)
	root = linked.Add(root, 3)
	root = linked.Add(root, 2)
	root = linked.Add(root, 1)

	n := linked.MidUp(root)
	fmt.Println(n.Val)
}

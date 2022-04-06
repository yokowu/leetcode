package main

import (
	"fmt"

	"github.com/yokowu/leetcode/linked"
)

func main() {
	root := linked.New(1)
	root = linked.Append(root, 2)
	root = linked.Append(root, 3)
	root = linked.Append(root, 4)
	root = linked.Append(root, 5)
	root = linked.Append(root, 6)

	linked.Show(root)

	fmt.Println("-----")
	root = linked.RotateRight(root, 2)
	linked.Show(root)
}

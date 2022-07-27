package main

import (
	"github.com/yokowu/leetcode/tree"
)

func main() {
	t := &tree.NodeTree{
		Left: &tree.NodeTree{
			Val: 10,
		},
		Right: &tree.NodeTree{
			Val: 11,
		},
		Val: 5,
	}
	tree.PostOrderTraversal(t)
}

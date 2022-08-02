package main

import (
	"github.com/yokowu/leetcode/tree"
)

func main() {
	t := &tree.TreeNode{
		Left: &tree.TreeNode{
			Val: 10,
		},
		Right: &tree.TreeNode{
			Val: 11,
		},
		Val: 5,
	}
	tree.PostOrderTraversal(t)
}

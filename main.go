package main

import (
	"fmt"

	"github.com/yokowu/leetcode/tree"
)

func main() {
	t := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 3,
			},
			Right: &tree.TreeNode{
				Val: 1,
			},
		},
		Right: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
	}
	res := tree.MaxDepth(t)
	fmt.Println(res)
}

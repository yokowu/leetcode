package main

import (
	"github.com/yokowu/leetcode/tree"
)

func main() {
	t := &tree.TreeNode{
		Left: &tree.TreeNode{
			Val:   2,
			Left:  &tree.TreeNode{Val: 4},
			Right: &tree.TreeNode{Val: 5},
		},
		Right: &tree.TreeNode{
			Val:   3,
			Left:  &tree.TreeNode{Val: 6},
			Right: &tree.TreeNode{Val: 7},
		},
		Val: 1,
	}
	tree.LevelTraversal(t)

	tree.Connect(&tree.Node{
		Left:  &tree.Node{Val: 2},
		Right: &tree.Node{Val: 3},
		Val:   1,
	})
}

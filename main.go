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
	t = tree.BuildTree2([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	tree.LevelTraversal(t)
}

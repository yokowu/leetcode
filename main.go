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
	t = tree.BuildTree2([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	tree.PreOrder(t)
}

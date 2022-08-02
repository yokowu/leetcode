package tree

import "fmt"

/*
满二叉树:
	如果一棵二叉树只有度为0的结点和度为2的结点，并且度为0的结点在同一层上，
	则这棵二叉树为满二叉树
            o
		   / \
		o       o
	   / \     / \
	  o   o   o   o
	 / \ / \ / \ / \
	o  o o o o o o o
深度为k, 有2^k-1个节点

完全二叉树:
	除了最底层节点，其余每层节点数都达到最大值
	最下面一层的节点都集中在该层最左边的若干位置
            o
		   / \
		o       o
	   / \     / \
	  o   o   o   o
	 / \ / \ / \
	o  o o o o o

二叉搜索树:
	若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值
	若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值
	它的左、右子树也分别为二叉排序树

平衡二叉搜索树:
	平衡二叉搜索树:又被称为AVL(Adelson-Velsky and Landis)树
	且具有以下性质:它是一棵空树或 它的左右两个子树的高度差的绝对值不超过1
	并且左右两个子树都是一棵平衡二叉树

递归三要素:
	1. 确定递归函数的参数和返回值: 确定哪些参数是递归的过程中需要处理的
		那么就在递归函数里加上这个参数，并且还要明确每次递归的返回值是什么
		进而确定递归函数的返回类型
	2. 确定终止条件
	3. 确定单层递归的逻辑: 确定每一层递归需要处理的信息
*/

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 前序遍历
func PreOrder(n *TreeNode) {
	if n == nil {
		return
	}
	var s Stack
	s.Push(n)
	for !s.IsEmpty() {
		n = s.Pop()
		fmt.Print(n.Val, " ")
		if n.Left != nil {
			s.Push(n.Left)
		}
		if n.Right != nil {
			s.Push(n.Right)
		}
	}
}

// 中序遍历
func InOrderTraversal(n *TreeNode) {
	var s Stack
	for n != nil || !s.IsEmpty() {
		for n != nil {
			s.Push(n)
			n = n.Left
		}

		if !s.IsEmpty() {
			n = s.Pop()
			fmt.Print(n.Val, " ")
			n = n.Right
		}
	}
}

// 后序遍历
func PostOrderTraversal(n *TreeNode) {
	var s Stack
	var last *TreeNode
	for n != nil || !s.IsEmpty() {
		for n != nil {
			s.Push(n)
			n = n.Left
		}

		n = s.Peek()
		if n.Right == nil || n.Right == last {
			fmt.Print(n.Val, " ")
			s.Pop()
			last = n
			n = nil
		} else {
			n = n.Right
		}
	}
}

// 层序遍历
func LevelTraversal(n *TreeNode) {
	var q Queue
	q.Enqueue(n)
	for !q.IsEmpty() {
		for i := 0; i < q.Len(); i++ {
			n = q.Dequeue()
			fmt.Print(n.Val, " ")
			if n.Left != nil {
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				q.Enqueue(n.Right)
			}
		}
	}
}

// 226. 翻转二叉树
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}

package tree

import (
	"fmt"
)

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
		if n.Right != nil {
			s.Push(n.Right)
		}
		if n.Left != nil {
			s.Push(n.Left)
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

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type NodeX struct {
	Val      int
	Children []*NodeX
}

// 116. 填充每个节点的下一个右侧节点指针
func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	connectTraversal(root.Left, root.Right)
	return root
}

func connectTraversal(n1, n2 *Node) {
	if n1 == nil || n2 == nil {
		return
	}

	n1.Next = n2
	connectTraversal(n1.Left, n1.Right)
	connectTraversal(n1.Right, n2.Left)
	connectTraversal(n2.Left, n2.Right)
}

// 114. 二叉树展开为链表
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}

	Flatten(root.Left)
	Flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}

// 654. 最大二叉树
func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := -1 << 31
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] > max {
			max = nums[j]
			i = j
		}
	}

	root := &TreeNode{Val: max}
	root.Left = ConstructMaximumBinaryTree(nums[:i])
	root.Right = ConstructMaximumBinaryTree(nums[i+1:])
	return root
}

// 105. 从前序与中序遍历序列构造二叉树
func BuildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == val {
			break
		}
	}

	root := &TreeNode{Val: val}
	size := len(inorder[:i]) + 1
	root.Left = BuildTree(preorder[1:size], inorder[:i])
	root.Right = BuildTree(preorder[size:], inorder[i+1:])
	return root
}

// 106. 从中序与后序遍历序列构造二叉树
func BuildTree2(inorder, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == val {
			break
		}
	}

	return &TreeNode{
		Val:   val,
		Left:  BuildTree2(inorder[:i], postorder[:i]),
		Right: BuildTree2(inorder[i+1:], postorder[i:len(postorder)-1]),
	}
}

// 889. 根据前序和后序遍历构造二叉树
func ConstructFromPrePost(preorder, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	if len(preorder) == 1 {
		return &TreeNode{Val: val}
	}

	nxt := preorder[1]
	i := 0
	for ; i < len(postorder); i++ {
		if postorder[i] == nxt {
			break
		}
	}
	size := i + 1
	return &TreeNode{
		Val:   val,
		Left:  ConstructFromPrePost(preorder[1:size+1], postorder[:size]),
		Right: ConstructFromPrePost(preorder[size+1:], postorder[size:len(postorder)-1]),
	}
}

// 107. 二叉树的层序遍历 II
func LeverlOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	que := make([]*TreeNode, 0)
	que = append(que, root)
	res := [][]int{}
	for len(que) > 0 {
		tmp := make([]int, 0)
		n := len(que)
		for i := 0; i < n; i++ {
			root = que[0]
			que = que[1:]
			tmp = append(tmp, root.Val)
			if root.Left != nil {
				que = append(que, root.Left)
			}
			if root.Right != nil {
				que = append(que, root.Right)
			}
		}

		if len(tmp) > 0 {
			newRes := [][]int{tmp}
			newRes = append(newRes, res...)
			res = newRes
		}
	}
	return res
}

// 199. 二叉树的右视图
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	que := make([]*TreeNode, 0)
	que = append(que, root)
	res := []int{}
	for len(que) > 0 {
		n := len(que)
		for i := 0; i < n; i++ {
			root = que[0]
			que = que[1:]

			if root.Left != nil {
				que = append(que, root.Left)
			}
			if root.Right != nil {
				que = append(que, root.Right)
			}
		}
		res = append(res, root.Val)
	}
	return res
}

// 637. 二叉树的层平均值
func AverageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	que := make([]*TreeNode, 0)
	que = append(que, root)
	res := []float64{}
	for len(que) > 0 {
		n := len(que)
		sum := 0
		for i := 0; i < n; i++ {
			root = que[0]
			que = que[1:]
			sum += root.Val

			if root.Left != nil {
				que = append(que, root.Left)
			}
			if root.Right != nil {
				que = append(que, root.Right)
			}
		}

		res = append(res, float64(sum)/float64(n))
		sum = 0
	}
	return res
}

// 429. N 叉树的层序遍历
func LevelOrder(root *NodeX) [][]int {
	if root == nil {
		return nil
	}
	que := make([]*NodeX, 0)
	que = append(que, root)
	res := make([][]int, 0)
	for len(que) > 0 {
		tmp := make([]int, len(que))
		tmpQ := make([]*NodeX, 0, len(que))
		for i, v := range que {
			tmp[i] = v.Val
			tmpQ = append(tmpQ, v.Children...)
		}
		que = tmpQ
		res = append(res, tmp)
	}
	return res
}

// 515. 在每个树行中找最大值
func LargestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var q Queue
	q.Enqueue(root)
	res := make([]int, 0)
	for !q.IsEmpty() {
		n := q.Len()
		max := q.Peek().Val
		for i := 0; i < n; i++ {
			root = q.Dequeue()
			if root.Val > max {
				max = root.Val
			}
			if root.Left != nil {
				q.Enqueue(root.Left)
			}
			if root.Right != nil {
				q.Enqueue(root.Right)
			}
		}
		res = append(res, max)
	}
	return res
}

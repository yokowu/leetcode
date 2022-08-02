package tree

type Queue struct {
	trees []*TreeNode
}

func (q *Queue) Enqueue(n *TreeNode) {
	q.trees = append(q.trees, n)
}

func (q *Queue) Dequeue() *TreeNode {
	n := q.trees[0]
	q.trees = q.trees[1:]
	return n
}

func (q *Queue) Len() int {
	return len(q.trees)
}

func (q *Queue) IsEmpty() bool {
	return len(q.trees) == 0
}

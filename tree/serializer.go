package tree

import (
	"bytes"
	"strconv"
	"strings"
)

func Serialize(root *TreeNode) string {
	buf := new(bytes.Buffer)
	serialize(root, buf)
	return buf.String()
}

func serialize(root *TreeNode, buf *bytes.Buffer) {
	if root == nil {
		buf.Write([]byte("#"))
		buf.Write([]byte(","))
		return
	}

	b := strconv.Itoa(root.Val)
	buf.Write([]byte(b))
	buf.Write([]byte(","))
	serialize(root.Left, buf)
	serialize(root.Right, buf)
}

func Deserialize(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}

	ss := strings.Split(s, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if ss[0] == "#" {
			ss = ss[1:]
			return nil
		}
		val, _ := strconv.Atoi(ss[0])
		ss = ss[1:]
		return &TreeNode{
			Val:   val,
			Left:  build(),
			Right: build(),
		}
	}
	return build()
}

package main

import "github.com/yokowu/leetcode/linked"

func main() {
	l := linked.New()
	l.Append(10)
	l.Append(5)
	l.Append(4)
	l.Append(3)
	l.Append(2)
	l.Append(1)
	l.Reverse()
	l.Show()
	l.Len()
}

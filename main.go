package main

import "github.com/yokowu/leetcode/hashset"

func main() {
	set := hashset.Constructor()
	set.Add(1)
	set.Add(10)
	set.Add(11)
	set.Add(1)
	set.Add(100)
	set.Remove(10)
	set.Show()
}

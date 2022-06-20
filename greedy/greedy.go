package greedy

import "sort"

func LongestPalindrome(s string) int {
	chars := make([]int, 128)
	for _, v := range s {
		chars[v]++
	}
	res := 0
	for _, v := range chars {
		res += v - (v % 2)
	}
	if res < len(s) {
		res++
	}
	return res
}

func FindContentChildren(g, s []int) int {
	sort.Ints(s)
	sort.Ints(g)

	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}

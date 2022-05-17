package greedy

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

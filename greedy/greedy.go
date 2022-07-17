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

// 分发饼干
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

// 摆动序列
func WiggleMaxLength(nums []int) int {
	result := 1
	preDiff, curDiff := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if curDiff > 0 && preDiff <= 0 || curDiff < 0 && preDiff >= 0 {
			result++
			preDiff = curDiff
		}
	}
	return result
}

// 最大子序和
func MaxSubArray(nums []int) int {
	result := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > result {
			result = count
		}
		if count <= 0 {
			count = 0
		}
	}
	return result
}

// 买卖股票的最佳时机2
func MaxProfit(prices []int) int {
	r := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			r += prices[i] - prices[i-1]
		}
	}
	return r
}

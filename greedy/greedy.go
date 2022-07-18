package greedy

import (
	"math"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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

// 跳跃游戏
func CanJump(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		if i+nums[i] > cover {
			cover = i + nums[i]
		}
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

// 跳跃游戏2
func Jump(nums []int) int {
	cur, next, ans := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > next {
			next = i + nums[i]
		}

		if cur == i {
			cur = next
			ans++
		}
	}
	return ans
}

func LargestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] *= -1
			k--
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}
	return sum(nums)
}

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func CanCompleteCircuit(gas, cost []int) int {
	sum, total, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if sum < 0 {
			sum = 0
			start = i + 1
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

func Candy(ratings []int) int {
	candy := make([]int, len(ratings))
	for i := 0; i < len(candy); i++ {
		candy[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candy[i] = max(candy[i], candy[i+1]+1)
		}
	}
	return sum(candy)
}

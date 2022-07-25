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

// k次取反后最大化的数组和
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

// 加油站
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

// 分发糖果
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

// 柠檬水找零
func LemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, v := range bills {
		switch v {
		case 5:
			five++
		case 10:
			if five <= 0 {
				return false
			}
			ten++
			five--
		case 20:
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// 用最少数量的箭引爆气球
func FindMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	r := 1
	prev := points[0]
	for i := 1; i < len(points); i++ {
		cur := points[i]
		if cur[0] > prev[1] {
			r++
			prev = cur
		}
	}
	return r
}

// 无重叠区间
func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	r := 1
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] >= prev[1] {
			r++
			prev = cur
		}
	}
	return len(intervals) - r
}

// 划分字母区间
func PartitionLabels(s string) []int {
	m := make(map[rune]int, 0)
	for i, v := range s {
		m[v] = i
	}

	res := make([]int, 0)
	left, right := 0, 0
	for i, v := range s {
		right = max(right, m[v])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] {
			res = append(res, prev)
			prev = cur
		} else {
			if cur[1] > prev[1] {
				prev[1] = cur[1]
			}
		}
	}
	res = append(res, prev)
	return res
}

// 714. 买卖股票的最佳时机含手续费
func MaxProfitFee(prices []int, fee int) int {
	min, r := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i] > min+fee {
			r += prices[i] - min - fee
			min = prices[i] - fee
		}
	}
	return r
}

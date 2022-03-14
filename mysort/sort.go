package mysort

import "math"

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func BubbleSort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
			}
		}
	}
}

func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		swap(nums, i, min)
	}
}

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0 && nums[j] > nums[j+1]; j-- {
			swap(nums, j, j+1)
		}
	}
}

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, L, R int) {
	if L == R {
		return
	}
	mid := ((R - L) / 2) + L
	mergeSort(nums, L, mid)
	mergeSort(nums, mid+1, R)
	merge(nums, L, mid, R)
}

func MergeSort2(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	mergeSize := 1
	n := len(nums)
	for mergeSize < n {
		l := 0
		for l < n {
			m := l + mergeSize - 1
			if m >= n {
				break
			}
			r := int(math.Min(float64(m+mergeSize), float64(n-1)))
			merge(nums, l, m, r)
			l = r + 1
		}
		if mergeSize > n/2 {
			break
		}
		mergeSize <<= 1
	}
}

func merge(nums []int, L, M, R int) {
	newer := make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := M + 1

	for p1 <= M && p2 <= R {
		if nums[p1] < nums[p2] {
			newer[i] = nums[p1]
			p1++
		} else {
			newer[i] = nums[p2]
			p2++
		}
		i++
	}

	for p1 <= M {
		newer[i] = nums[p1]
		p1++
		i++
	}

	for p2 <= R {
		newer[i] = nums[p2]
		p2++
		i++
	}

	for i := 0; i < len(newer); i++ {
		nums[L+i] = newer[i]
	}
}

// Partition 给定一个数组nums, 要求把数组中小于target的数放在左边, 大于target的数放在右边
func Partition(nums []int, target int) {
	min := 0
	i := 0
	for i < len(nums) {
		if nums[i] <= target {
			swap(nums, i, min)
			min++
		}
		i++
	}

}

// Partition2 给定一个数组nums,
// 要求把数组中小于target的数放左边,
// 等于target的数放中间
// 大于target的放右边
func Partition2(nums []int, target int) {
	min := 0
	max := len(nums) - 1
	i := 0
	for i <= max {
		if nums[i] < target {
			swap(nums, i, min)
			min++
			i++
		} else if nums[i] == target {
			i++
		} else if nums[i] > target {
			swap(nums, i, max)
			max--
		}
	}
}

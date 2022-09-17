package mysort

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MergeSort2(nums []int) {
	if len(nums) < 2 {
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
			r := min(m+mergeSize, n-1)
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

func HeapSort(nums []int) {
	n := len(nums)
	// build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		swap(nums, 0, i)
		n--
		heapify(nums, 0, n)
	}
}

func heapify(nums []int, i, n int) {
	max := i
	l := i*2 + 1
	r := i*2 + 2
	if l < n && nums[l] < nums[max] {
		max = l
	}
	if r < n && nums[r] < nums[max] {
		max = r
	}

	if i != max {
		swap(nums, max, i)
		heapify(nums, max, n)
	}
}

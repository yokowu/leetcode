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

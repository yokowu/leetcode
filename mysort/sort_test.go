package mysort

import (
	"math/rand"
	"sort"
	"testing"
)

type Sint []int

func (a Sint) Len() int           { return len(a) }
func (a Sint) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Sint) Less(i, j int) bool { return a[i] < a[j] }

func genRandNums(n int) Sint {
	nums := make(Sint, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(100)
	}
	return nums
}

func TestBubbleSort(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		nums := genRandNums(10)
		BubbleSort(nums)

		if !sort.IsSorted(nums) {
			t.Errorf("%+v", nums)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		nums := genRandNums(10)
		SelectionSort(nums)

		if !sort.IsSorted(nums) {
			t.Errorf("%+v", nums)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for i := 0; i < 1; i++ {
		nums := genRandNums(10)
		InsertionSort(nums)

		if !sort.IsSorted(nums) {
			t.Errorf("%+v", nums)
		}
	}
}

func BenchmarkBubble(b *testing.B) {
	nums := genRandNums(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleSort(nums)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	nums := genRandNums(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SelectionSort(nums)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	nums := genRandNums(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		InsertionSort(nums)
	}
}

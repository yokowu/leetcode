package heap

import (
	"fmt"
)

type Heap struct {
	size   int
	len    int
	values []int
}

func New(len int) *Heap {
	return &Heap{
		len:    len,
		values: make([]int, len),
	}
}

func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func (h *Heap) Push(val int) {
	if h.size == h.len {
		fmt.Printf("heap full val: %d\n", val)
		return
	}

	h.values[h.size] = val
	h.insert(h.size)
	h.size++
}

func (h *Heap) insert(i int) {
	for h.values[i] > h.values[(i-1)/2] {
		h.swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func (h *Heap) swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *Heap) Show() {
	fmt.Printf("%+v size: %d \n", h.values, h.size)
}
